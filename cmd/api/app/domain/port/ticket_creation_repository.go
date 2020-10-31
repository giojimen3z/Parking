package port

import "github.com/Parking/cmd/api/app/domain/model"

// TicketCreationRepository interface to connect bike access implementation
type TicketCreationRepository interface {
	// SaveTicket post the ticket into DBA
	SaveTicket(ticket model.Ticket) (err error)
}
