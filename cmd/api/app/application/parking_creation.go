package application

import (
	"fmt"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/pkg/logger"
)

const (
	errorServiceParking = "error getting information from service Parking Creation for Parking %v"
)

// ParkingCreationApplication is the initial flow entry to granted the create  the parking
type ParkingCreationApplication interface {
	// Handler is the input for create the parking
	Handler(parking model.Parking) (err error)
}

type ParkingCreation struct {
	ParkingCreationService service.ParkingCreationServicePort
}

func (parkingCreation *ParkingCreation) Handler(parking model.Parking) (err error) {

	err = parkingCreation.ParkingCreationService.ParkingCreation(parking)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceParking, parking.ParkingName), err)
		return err
	}
	return err
}
