package application

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/service"
)

// AccessBikeApplication is the initial flow entry to granted the access to the parking
type AccessBikeApplication interface {
	// Handler is the input for access to the parking
	Handler(bike model.Bike) (err error)
}

// BikeAccess represents  the use case  to Save the bike into DBA
type BikeAccess struct {
	BikeAccessService service.BikeAccessServicePort
}

// Handler execute the BikeAccess for save  the bike into DBA
func (accessBike *BikeAccess) Handler(bike model.Bike) error {
	return accessBike.BikeAccessService.BikeAccess(bike)

}
