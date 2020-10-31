package service

import (
	"errors"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/errorApi/logger"
)

const (
	errorParkingListRepository = "Error Getting the parking list from repository"
)

type ParkingListServicePort interface {
	// ParkingList Get from Repository all parking
	ParkingList() (parkingList []model.Parking, err error)
}

type ParkingListService struct {
	ParkingListRepository port.ParkingListRepository
}

// BikeAccess process the information
func (bikeAccessService *ParkingListService) ParkingList() (parkingLot []model.Parking, err error) {

	parkingLot, err = bikeAccessService.ParkingListRepository.ListParking()

	if err != nil {
		err = errors.New(errorParkingListRepository)
		logger.Error(errorParkingListRepository, err)
		return parkingLot, err
	}

	return parkingLot, err
}
