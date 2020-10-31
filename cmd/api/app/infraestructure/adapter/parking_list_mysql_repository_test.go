package adapter_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/cmd/api/app/infraestructure/adapter"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/stretchr/testify/assert"
)

const (
	ParkingListQuery = "SELECT (.+) FROM parking "
)

func setUpParkingListRepository() (parkingListRepository port.ParkingListRepository, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	parkingListRepository = &adapter.ParkingListMysqlRepository{
		ReadConnectionClient: db,
	}
	return
}

func TestWhenGetParkingListOkThenReturnParkingListAndNilError(t *testing.T) {
	parkingList := []model.Parking{builder.NewParkingDataBuilder().Build()}
	parking := builder.NewParkingDataBuilder().Build()
	repository, dbMock := setUpParkingListRepository()
	rows := sqlmock.NewRows([]string{"id", "name", "address", "owner"}).AddRow(parking.ParkingId, parking.ParkingName, parking.ParkingAddress, parking.ParkingOwner)
	dbMock.ExpectQuery(ParkingListQuery).WillReturnRows(rows)

	parkingLot, errorResult := repository.ListParking()

	assert.Nil(t, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
	assert.Equal(t, parkingList, parkingLot)
}

func TestWhenTryGetParkingListThenReturnParkingListEmptyAndError(t *testing.T) {
	errorExpected := "some type of parameters is not correct"
	repository, dbMock := setUpParkingListRepository()
	dbMock.ExpectQuery(ParkingListQuery).WillReturnError(errors.New(errorExpected))

	parkingLot, errorResult := repository.ListParking()

	assert.NotNil(t, errorResult)
	assert.Nil(t, parkingLot)
	assert.Equal(t, errorExpected, errorResult.Error())
	assert.Nil(t, dbMock.ExpectationsWereMet())

}
func TestWhenTryGetParkingListQueryScanIsWrongThenFailed(t *testing.T) {

	parking := builder.NewParkingDataBuilder().Build()
	errorExpected := "sql: expected 3 destination arguments in Scan, not 4"
	repository, dbMock := setUpParkingListRepository()
	rows := sqlmock.NewRows([]string{"id", "name", "address"}).AddRow(parking.ParkingId, parking.ParkingName, parking.ParkingAddress)
	dbMock.ExpectQuery(ParkingListQuery).WillReturnRows(rows)

	parkingLot, errorResult := repository.ListParking()

	assert.NotNil(t, errorResult)
	assert.Nil(t, parkingLot)
	assert.Equal(t, errorExpected, errorResult.Error())
	assert.Nil(t, dbMock.ExpectationsWereMet())

}
