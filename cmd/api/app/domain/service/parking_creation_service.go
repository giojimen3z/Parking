package service

import (
	"errors"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/pkg/logger"
)

type ParkingCreationServicePort interface {
	// ParkingCreation Send to Repository the parking requested
	ParkingCreation(parking model.Parking) (err error)
}

type ParkingCreationService struct {
	ParkingCreationRepository port.ParkingCreationRepository
}

func (parkingCreationService *ParkingCreationService) ParkingCreation(parking model.Parking) (err error) {

	err = parkingCreationService.ParkingCreationRepository.SaveParking(parking)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
