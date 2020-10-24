package model

type Bike struct {
	BikeId       int64  `json:"bike_id"`
	SerialNumber string `json:"serial_number"`
	Brand        string `json:"brand"`
	Color        string `json:"color"`
}
