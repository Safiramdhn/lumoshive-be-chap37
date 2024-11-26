package router

import (
	"api-service-shipping/controller"

	"github.com/gin-gonic/gin"
)

// APIRouter configures the API routes for the shipping service
func APIRouter(router *gin.Engine, ctl controller.MainController) {
	// Shipping-related routes
	shipping := router.Group("/shipping")
	{
		shipping.GET("/list", ctl.ShippingController.GetAllShippingController)
		shipping.GET("/:id", ctl.ShippingController.GetShippingByIdController)
		shipping.GET("/cost/:id/:quantity/:origin_longlat/:destination_longlat", ctl.ShippingController.GetShippingCostController)
	}
}
