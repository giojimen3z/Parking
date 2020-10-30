package port

// TicketCreationRepository interface to connect bike access implementation
type TicketCreationRepository interface {
	// SaveBike post the Bike into DBA
	SaveTicket(serialNumber string, enterDate string) (err error)
}
