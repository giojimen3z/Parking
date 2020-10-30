package application

import "github.com/Parking/cmd/api/app/domain/model"

// ParkingCreationApplication is the initial flow entry to granted the create  the parking
type ParkingCreationApplication interface {
	// Handler is the input for create the parking
	Handler(parking model.Parking) (err error)
}
