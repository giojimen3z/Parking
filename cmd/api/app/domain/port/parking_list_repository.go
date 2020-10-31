package port

import "github.com/Parking/cmd/api/app/domain/model"

// ParkingListRepository interface to connect with the parking list implementation
type ParkingListRepository interface {
	// ListParking get the parking from DBA
	ListParking() (parkingLot []model.Parking, err error)
}
