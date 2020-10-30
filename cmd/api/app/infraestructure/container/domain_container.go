package container

import (
	"github.com/Parking/cmd/api/app/domain/service"
)

func getBikeAccessService() service.BikeAccessServicePort {
	return &service.BikeAccessService{
		BikeAccessRepository: getBikeAccessRepository(),
	}
}

func getTicketService() service.TicketCreationServicePort {
	return &service.TicketCreationService{
		TicketCreationRepository: getTicketRepository(),
	}
}
