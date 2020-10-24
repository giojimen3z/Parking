package model

import "time"

type Ticket struct {
	TicketId     int64     `json:"ticket_id"`
	TotalTime    time.Time `json:"total_time"`
	PaymentTotal float64   `json:"Payment_total"`
}
