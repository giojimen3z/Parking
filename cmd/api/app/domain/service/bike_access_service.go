package service

import (
	"errors"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/port"
)

const (
	errorRepository = "error getting repository information"
)

type BikeAccessServicePort interface {
	// BikeAccess Send to repository  the Bike requested
	BikeAccess(bike model.Bike) (err error)
}

type BikeAccessService struct {
	BikeAccessRepository port.BikeAccessRepository
}

// BikeAccess process the information
func (bikeAccessService *BikeAccessService) BikeAccess(bike model.Bike) (err error) {

	err = bikeAccessService.BikeAccessRepository.SaveBike(bike)

	if err != nil {
		err = errors.New(errorRepository)
		return err
	}

	return err
}
