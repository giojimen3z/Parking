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
	queryToSaveParking = `INSERT INTO parking (name,address,owner)VALUES(?,?,?)`
	errorSavingParking = "an error occurred save parking: %s"
	transactionError   = ""
)

// ParkingCreationMysqlRepository represent the mysql repository
type ParkingCreationMysqlRepository struct {
	WriteClient *sql.DB
}

//SaveParking is a function to initialize connection to the DB, take control of the transaction before returning something and send to save.
func (parkingCreationMysqlRepository *ParkingCreationMysqlRepository) SaveParking(parking model.Parking) (err error) {
	var tx *sql.Tx

	defer func() {
		config.CloseConnections(err, tx, nil, nil)
	}()

	tx, err = parkingCreationMysqlRepository.WriteClient.Begin()
	if err != nil {
		errMsg := fmt.Sprintf(errorSavingParking, parking.ParkingName)
		logger.Error(errMsg, err)
		return exception.InternalServerError{ErrMessage: errMsg}
	}
	_, err = parkingCreationMysqlRepository.WriteClient.Exec(queryToSaveParking,
		parking.ParkingName,
		parking.ParkingAddress,
		parking.ParkingOwner)
	if err != nil {
		errMsg := errors.New(errorParameter)
		logger.Error(errMsg.Error(), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}

	return err
}
