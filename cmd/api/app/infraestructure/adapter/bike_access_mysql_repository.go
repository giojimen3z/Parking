package adapter

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/infraestructure/config"
	"github.com/Parking/errorApi/logger"
)

const (
	queryToSaveBike = `INSERT INTO bike (serial_number,brand,color)VALUES(?,?,?)`
	errorParameter  = "the type of parameter is not correct"
	errorSavingBike = "an error occurred save bike with serial number %s"

)

// BikeAccessMysqlRepository represent the mysql repository
type BikeAccessMysqlRepository struct {
	WriteClient *sql.DB
}

//SaveBike is a function to initialize connection to the DB, take control of the transaction before returning something and send to save.
func (bikeAccessMysqlRepository *BikeAccessMysqlRepository) SaveBike(bike model.Bike) (err error) {
	var tx *sql.Tx

	defer func() {
		config.CloseConnections(err, tx, nil, nil)
	}()

	tx, err = bikeAccessMysqlRepository.WriteClient.Begin()
	if err != nil {
		errMsg := fmt.Sprintf(errorSavingBike, bike.SerialNumber)
		logger.Error(errMsg, err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}
	_, err = bikeAccessMysqlRepository.WriteClient.Exec(queryToSaveBike,
		bike.SerialNumber,
		bike.Brand,
		bike.Color)
	if err != nil {
		errMsg := errors.New(errorParameter)
		logger.Error(errMsg.Error(), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}

	return err
}
