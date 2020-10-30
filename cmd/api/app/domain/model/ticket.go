package model

import "time"

type Ticket struct {
	TicketId     int64     `json:"ticket_id"`
	TicketNumber string    `json:"ticket_number"`
	EnterDate    time.Time `json:"enter_date"`
	ExitDate     time.Time `json:"exit_date"`
	TotalTime    string    `json:"total_time"`
	PaymentTotal float64   `json:"payment_total"`
}
