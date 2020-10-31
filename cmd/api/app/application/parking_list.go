package application

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/errorApi/logger"
)

const (
	errorServiceParkingList = "error getting information from service Parking List"
)

// ParkingListApplication is the initial flow entry to granted the get to the parking
type ParkingListApplication interface {
	// Handler is the input for get the parking
	Handler() (parkingLots []model.Parking, err error)
}

type ParkingList struct {
	ParkingListService service.ParkingListServicePort
}

func (parkingList *ParkingList) Handler() (parkingLots []model.Parking, err error) {

	parkingLots, err = parkingList.ParkingListService.ParkingList()

	if err != nil {
		logger.Error(errorServiceParkingList, err)
		return parkingLots, err
	}
	return parkingLots, err
}
