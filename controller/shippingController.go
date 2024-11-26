package controller

import (
	"api-service-shipping/model"
	"api-service-shipping/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShippingController struct {
	Service service.ShippingService
}

func NewShippingController(db *sql.DB) ShippingController {
	return ShippingController{Service: service.NewShippingService(db)}
}

func (c ShippingController) GetAllShippingController(ctx *gin.Context) {
	shippingList, err := c.Service.GetAllShippings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting shipping list": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shippingList})
}

func (c ShippingController) GetShippingByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	shippingId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	shipping, err := c.Service.GetShippingById(shippingId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error getting shipping": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shipping})
}

func (c ShippingController) GetShippingCostController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	originLatLong := ctx.Param("origin_longlat")
	destinationLatLong := ctx.Param("destination_longlat")
	quantity := ctx.Param("quantity")
	if originLatLong == "" || destinationLatLong == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "origin and destination latitude/longitude are required"})
		return
	}
	shippingId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ItemQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
		return
	}

	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s?overview=false", originLatLong, destinationLatLong)
	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting shipping cost": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error parsing response": err.Error()})
		return
	}

	input := model.ShippingCostRequest{
		ShippingID: shippingId,
		Quantity:   ItemQuantity,
	}

	shippingCost, err := c.Service.CalculateShippingCost(result, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error calculating shipping cost": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shippingCost})
}
