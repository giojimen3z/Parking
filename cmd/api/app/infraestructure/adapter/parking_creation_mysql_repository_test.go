package adapter_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/cmd/api/app/infraestructure/adapter"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/stretchr/testify/assert"
)

const (
	insertParkingQuery = "INSERT INTO parking "
)

func setUpParkingRepository() (parkingCreationRepository port.ParkingCreationRepository, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	parkingCreationRepository = &adapter.ParkingCreationMysqlRepository{
		WriteClient: db,
	}
	return
}
func TestWhenSaveParkingIsOkThenReturnNil(t *testing.T) {
	parking := builder.NewParkingDataBuilder().Build()
	repository, dbMock := setUpParkingRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertParkingQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	errorResult := repository.SaveParking(parking)

	assert.Nil(t, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveParkingTransactionFailThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error occurred save parking: PArkAutosBosa"
	parking := builder.NewParkingDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpParkingRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertParkingQuery).WillReturnError(errorOnUpdate)

	errorResult := repository.SaveParking(parking)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveParkingTransactionBeginErrorThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error occurred save parking: PArkAutosBosa"
	parking := builder.NewParkingDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpParkingRepository()

	dbMock.ExpectBegin().WillReturnError(errorOnUpdate)

	errorResult := repository.SaveParking(parking)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
