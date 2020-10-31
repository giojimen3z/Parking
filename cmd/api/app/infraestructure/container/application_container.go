package container

import (
	"github.com/Parking/cmd/api/app/application"
)

func GetBikeAccessApplication() application.AccessBikeApplication {
	return &application.BikeAccess{
		BikeAccessService:     getBikeAccessService(),
		TicketCreationService: getTicketService(),
	}
}
func GetParkingAccessApplication() application.ParkingCreationApplication {
	return &application.ParkingCreation{ParkingCreationService: getParkingService()}
}
func GetParkingListApplication() application.ParkingListApplication {
	return &application.ParkingList{ParkingListService: getParkingListService()}
}
