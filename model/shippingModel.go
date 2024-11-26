package model

type Shipping struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

type ShippingCostRequest struct {
	ShippingID         int    `json:"id_shipping"`
	Quantity           int    `json:"quantity_barang"`
	OriginLatLong      string `json:"origin_latlong"`
	DestinationLatLong string `json:"destination_latlong"`
}

type ShippingCostResponse struct {
	Distance float64 `json:"distance"`
	Cost     float64 `json:"cost"`
}
