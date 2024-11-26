package controller

import (
	"database/sql"
)

type MainController struct {
	ShippingController ShippingController
}

func NewMainController(db *sql.DB) MainController {
	return MainController{
		ShippingController: NewShippingController(db),
	}
}
