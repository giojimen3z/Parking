package model

type Parking struct {
	ParkingId      int64  `json:"parking_id"`
	ParkingName    string `json:"parking_name"`
	ParkingAddress string `json:"parking_address"`
	ParkingOwner   string `json:"parking_owner"`
}
