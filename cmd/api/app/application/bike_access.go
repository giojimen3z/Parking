package application

import (
	"fmt"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/errorApi/logger"
)

const (
	errorService        = "error getting information from service Access Bike"
	ErrorTicketCreation = "error creating ticket by Bike with serial number : %v"
)

// AccessBikeApplication is the initial flow entry to granted the access to the parking
type AccessBikeApplication interface {
	// Handler is the input for access to the parking
	Handler(bike model.Bike) (err error)
}

// BikeAccess represents  the use case  to Save the bike into DBA
type BikeAccess struct {
	BikeAccessService     service.BikeAccessServicePort
	TicketCreationService service.TicketCreationServicePort
}

// Handler execute the BikeAccess for save  the bike into DBA
func (accessBike *BikeAccess) Handler(bike model.Bike) (err error) {

	err = accessBike.BikeAccessService.BikeAccess(bike)
	if err != nil {
		logger.Error(errorService, err)
		return err
	}

	err = accessBike.TicketCreationService.TicketCreation(bike.SerialNumber)
	if err != nil {
		logger.Error(fmt.Sprintf(ErrorTicketCreation, bike.SerialNumber), err)
		return err
	}
	return err

}
