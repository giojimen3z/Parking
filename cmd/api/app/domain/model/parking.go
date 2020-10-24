package model

import "time"

type Parking struct {
	ParkingId        int64     `json:"parking_id"`
	ParkingEnterDate time.Time `json:"parking_enter_date"`
	ParkingExitDate  time.Time `json:"Parking_exit_date"`
}
