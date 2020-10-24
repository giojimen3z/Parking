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
	insertQueryMatcher = "INSERT INTO bike "
)

func setUpClaimRepository() (bikeAccessRepository port.BikeAccessRepository, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	bikeAccessRepository = &adapter.BikeAccessMysqlRepository{
		WriteClient: db,
	}
	return
}
func TestWhenSaveBikeIsOkThenReturnNil(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	repository, dbMock := setUpClaimRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertQueryMatcher).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	errorResult := repository.SaveBike(bike)

	assert.Nil(t, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveTransactionFailsThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error happened when execute the transaction"
	bike := builder.NewBikeDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpClaimRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertQueryMatcher).WillReturnError(errorOnUpdate)

	errorResult := repository.SaveBike(bike)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveTransactionBeginErrorThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error happened when initializing the transaction"
	bike := builder.NewBikeDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpClaimRepository()

	dbMock.ExpectBegin().WillReturnError(errorOnUpdate)

	errorResult := repository.SaveBike(bike)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
