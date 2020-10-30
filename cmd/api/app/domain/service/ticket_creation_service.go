package service

import (
	"errors"
	"time"

	"github.com/Parking/cmd/api/app/domain/port"
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

	enterDate:= time.Now().UTC().Format(time.RFC3339)
	err = ticketCreationService.TicketCreationRepository.SaveTicket(serialNumber,enterDate)

	if err != nil {
		err = errors.New(errorRepository)
		return err
	}

	return err
}