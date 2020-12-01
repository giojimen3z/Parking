package adapter

import (
	"database/sql"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/pkg/logger"
)

const (
	queryGetAllParking  = `SELECT * FROM parking`
	errorSyntaxQuery    = "query syntax error"
	errorScanningFields = "error when scanning the information"
)

type ParkingListMysqlRepository struct {
	ReadConnectionClient *sql.DB
}

func (parkingListMysqlRepository *ParkingListMysqlRepository) ListParking() (parkingList []model.Parking, err error) {

	var rowsParking *sql.Rows

	rowsParking, err = parkingListMysqlRepository.ReadConnectionClient.Query(queryGetAllParking)
	if err != nil {
		logger.Error(errorSyntaxQuery, err)
		return
	}
	defer rowsParking.Close()

	for rowsParking.Next() {
		parking := model.Parking{}
		err = rowsParking.Scan(
			&parking.ParkingId,
			&parking.ParkingName,
			&parking.ParkingAddress,
			&parking.ParkingOwner,
		)
		if err != nil {
			logger.Error(errorScanningFields, err)
			return
		}
		parkingList = append(parkingList, parking)
	}
	err = rowsParking.Err()

	return parkingList, err
}
