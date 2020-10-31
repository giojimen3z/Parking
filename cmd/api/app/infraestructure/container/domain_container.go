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
func getParkingService() service.ParkingCreationServicePort {
	return &service.ParkingCreationService{
		ParkingCreationRepository: getParkingCreationRepository(),
	}
}

func getParkingListService() service.ParkingListServicePort {
	return &service.ParkingListService{
		ParkingListRepository: getParkingListRepository(),
	}
}
