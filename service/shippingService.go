package service

import (
	"api-service-shipping/model"
	"api-service-shipping/repository"
	"database/sql"
	"errors"
)

type ShippingService interface {
	GetAllShippings() ([]model.Shipping, error)
	GetShippingById(id int) (*model.Shipping, error)
	CalculateShippingCost(apiRes map[string]interface{}, costReq model.ShippingCostRequest) (*model.ShippingCostResponse, error)
	CalculateCost(distance float64, quantity int) float64
}

type shippingService struct {
	Repo repository.ShippingRepository
}

func NewShippingService(db *sql.DB) ShippingService {
	return &shippingService{Repo: repository.NewShippingRepository(db)}
}

// GetAllShippings implements ShippingService.
func (s *shippingService) GetAllShippings() ([]model.Shipping, error) {
	return s.Repo.GetAll()
}

// GetShippingById implements ShippingService.
func (s *shippingService) GetShippingById(id int) (*model.Shipping, error) {
	if id == 0 {
		return nil, errors.New("id cannot be 0")
	}
	return s.Repo.GetByID(id)
}

func (s *shippingService) CalculateShippingCost(apiRes map[string]interface{}, costReq model.ShippingCostRequest) (*model.ShippingCostResponse, error) {
	routes := apiRes["routes"].([]interface{})
	if len(routes) == 0 {
		return nil, errors.New("no routes found")
	}

	distance := routes[0].(map[string]interface{})["distance"].(float64)
	shippingData, err := s.GetShippingById(costReq.ShippingID)
	if err != nil {
		return nil, err
	}
	// Convert from meters to kilometers
	distanceInKM := distance / 1000
	shippingCost := distanceInKM * shippingData.Rate
	shippingCost += s.CalculateCost(distanceInKM, costReq.Quantity)

	result := model.ShippingCostResponse{
		Distance: distanceInKM,
		Cost:     shippingCost,
	}

	return &result, nil
}

func (s *shippingService) CalculateCost(distance float64, quantity int) float64 {
	var costPerKm int
	if quantity < 2 {
		costPerKm = 2000
	} else {
		costPerKm = 4000
	}
	return distance * float64(costPerKm)
}
