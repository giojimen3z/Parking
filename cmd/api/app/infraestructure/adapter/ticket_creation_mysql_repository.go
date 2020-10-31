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
	errorSavingTicket = "an error occurred save bike with serial number %s"
	errorSavingDate   = "an error occurred with the format date"
	queryToSaveTicket = `INSERT INTO ticket (bike_serial,enter_date)VALUES(?,?)`
)

// TicketCreationMysqlRepository represent the mysql repository
type TicketCreationMysqlRepository struct {
	WriteClient *sql.DB
}

//SaveTicket is a function to initialize connection to the DB, take control of the transaction before returning something and send to save.
func (ticketCreationMysqlRepository *TicketCreationMysqlRepository) SaveTicket(ticket model.Ticket) (err error) {
	var tx *sql.Tx

	defer func() {
		config.CloseConnections(err, tx, nil, nil)
	}()


	tx, err = ticketCreationMysqlRepository.WriteClient.Begin()
	if err != nil {
		errMsg := fmt.Sprintf(errorSavingTicket, ticket.SerialNumber)
		logger.Error(errMsg, err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}
	_, err = ticketCreationMysqlRepository.WriteClient.Exec(queryToSaveTicket,
		ticket.SerialNumber,
		ticket.EnterDate)

	if err != nil {
		errMsg := errors.New(errorParameter)
		logger.Error(errMsg.Error(), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}

	return err
}
