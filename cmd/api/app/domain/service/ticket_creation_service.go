package service

import (
	"errors"
	"time"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/port"
	"github.com/Parking/pkg/logger"
)

type TicketCreationServicePort interface {
	// TicketCreation Create into DBA the Ticket
	TicketCreation(serialNumber string) (err error)
}

type TicketCreationService struct {
	TicketCreationRepository port.TicketCreationRepository
}

// TicketCreation process the information
func (ticketCreationService *TicketCreationService) TicketCreation(serialNumber string) (err error) {

	ticket := model.Ticket{}
	enterDate := time.Now().UTC().Format(time.RFC3339)
	ticket.EnterDate, _ = time.Parse(time.RFC3339, enterDate)
	ticket.SerialNumber = serialNumber

	err = ticketCreationService.TicketCreationRepository.SaveTicket(ticket)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
