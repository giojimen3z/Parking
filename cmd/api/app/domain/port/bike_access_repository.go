package port

import "github.com/Parking/cmd/api/app/domain/model"

// BikeAccessRepository interface to connect bike access implementation
type BikeAccessRepository interface {
	// SaveBike post the Bike into DBA
	SaveBike(bike model.Bike) (err error)
}
