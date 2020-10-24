package container

import (
	"database/sql"

	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/cmd/api/app/infraestructure/adapter"
	"github.com/Parking/cmd/api/app/infraestructure/config"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
)

func GetContentController() *controller.BikeAccessController {
	return &controller.BikeAccessController{BikeAccessApplication: GetBikeAccessApplication()}
}

func getBikeAccessRepository() port.BikeAccessRepository{
	return &adapter.BikeAccessMysqlRepository{WriteClient: getWriteConnectionClient()}
}
func getWriteConnectionClient() *sql.DB {
	conn, _ := config.GetWriteConnection()
	return conn
}
