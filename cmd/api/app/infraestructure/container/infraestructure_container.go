package container

import (
	"database/sql"

	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/cmd/api/app/infraestructure/adapter"
	"github.com/Parking/cmd/api/app/infraestructure/config"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
)

func GetBikeAccessController() *controller.BikeAccessController {
	return &controller.BikeAccessController{BikeAccessApplication: GetBikeAccessApplication()}
}

func GetParkingCreationController() *controller.ParkingCreationController {
	return &controller.ParkingCreationController{ParkingCreationApplication: GetParkingAccessApplication()}
}
func getBikeAccessRepository() port.BikeAccessRepository {
	return &adapter.BikeAccessMysqlRepository{WriteClient: getWriteConnectionClient()}
}

func GetParkingListController() *controller.ParkingListController {
	return &controller.ParkingListController{ParkingListApplication: GetParkingListApplication()}
}
func getTicketRepository() port.TicketCreationRepository {
	return &adapter.TicketCreationMysqlRepository{WriteClient: getWriteConnectionClient()}
}

func getParkingCreationRepository() port.ParkingCreationRepository {
	return &adapter.ParkingCreationMysqlRepository{WriteClient: getWriteConnectionClient()}
}

func getParkingListRepository() port.ParkingListRepository {
	return &adapter.ParkingListMysqlRepository{ReadConnectionClient: getReadConnectionClient()}
}

func getWriteConnectionClient() *sql.DB {
	conn, _ := config.GetWriteConnection()
	return conn
}

func getReadConnectionClient() *sql.DB {
	conn, _ := config.GetReadConnection()
	return conn
}
