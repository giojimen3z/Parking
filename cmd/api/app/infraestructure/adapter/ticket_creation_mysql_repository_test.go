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
	insertQueryTicket = "INSERT INTO ticket "
)

func setUpTicketRepository() (ticketCreationRepository port.TicketCreationRepository, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	ticketCreationRepository = &adapter.TicketCreationMysqlRepository{
		WriteClient: db,
	}
	return
}
func TestWhenSaveTicketIsOkThenReturnNil(t *testing.T) {
	ticket := builder.NewTicketDataBuilder().Build()
	repository, dbMock := setUpTicketRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertQueryTicket).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	errorResult := repository.SaveTicket(ticket)

	assert.Nil(t, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveTicketTransactionFailThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error happened when execute the transaction"
	ticket := builder.NewTicketDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpTicketRepository()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(insertQueryTicket).WillReturnError(errorOnUpdate)

	errorResult := repository.SaveTicket(ticket)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
func TestWhenSaveTicketTransactionBeginErrorThenReturnError(t *testing.T) {
	transactionErrorMessage := "an error happened when initializing the transaction"
	ticket := builder.NewTicketDataBuilder().Build()
	errorOnUpdate := exception.InternalServerError{ErrMessage: transactionErrorMessage}
	repository, dbMock := setUpTicketRepository()
	dbMock.ExpectBegin().WillReturnError(errorOnUpdate)

	errorResult := repository.SaveTicket(ticket)

	assert.NotNil(t, errorResult)
	assert.Equal(t, errorOnUpdate, errorResult)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
