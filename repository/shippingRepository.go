package repository

import (
	"api-service-shipping/model"
	"database/sql"
)

type ShippingRepository interface {
	GetAll() ([]model.Shipping, error)
	GetByID(id int) (*model.Shipping, error)
}

type shippingRepository struct {
	DB *sql.DB
}

func NewShippingRepository(db *sql.DB) ShippingRepository {
	return &shippingRepository{DB: db}
}

// GetAll implements ShippingRepository.
func (repo *shippingRepository) GetAll() ([]model.Shipping, error) {
	sqlStatement := "SELECT id, name FROM shipping_services"

	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shippings []model.Shipping
	for rows.Next() {
		var shipping model.Shipping
		err := rows.Scan(&shipping.ID, &shipping.Name)
		if err != nil {
			return nil, err
		}
		shippings = append(shippings, shipping)
	}
	return shippings, nil
}

// GetByID implements ShippingRepository.
func (repo *shippingRepository) GetByID(id int) (*model.Shipping, error) {
	sqlStatement := "SELECT id, name, cost_rate FROM shipping_services WHERE id=$1"

	var shipping model.Shipping
	err := repo.DB.QueryRow(sqlStatement, id).Scan(&shipping.ID, &shipping.Name, &shipping.Rate)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &shipping, nil
}
