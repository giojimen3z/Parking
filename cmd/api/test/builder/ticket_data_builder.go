package builder

import (
	"time"

	"github.com/Parking/cmd/api/app/domain/model"
)

type TicketDataBuilder struct {
	ticketId     int64
	ticketNumber string
	enterDate    time.Time
	exitDate     time.Time
	totalTime    string
	paymentTotal int64
	serialNumber string
}

func NewTicketDataBuilder() *TicketDataBuilder {
	enterDateString := time.Now().UTC().Format(time.RFC3339)
	enterDate, _ := time.Parse(time.RFC3339, enterDateString)
	exitDate, _ := time.Parse(time.RFC3339, "2020-10-29 23:40:47")
	return &TicketDataBuilder{
		ticketId:     1,
		ticketNumber: "N° 1",
		enterDate:    enterDate,
		exitDate:     exitDate,
		totalTime:    "2 hrs",
		paymentTotal: 1400,
		serialNumber: "MR145987D12",
	}
}
func (builder *TicketDataBuilder) WithTickedID(ticketId int64) *TicketDataBuilder {
	builder.ticketId = ticketId
	return builder
}
func (builder *TicketDataBuilder) WithTicketNumber(ticketNumber string) *TicketDataBuilder {
	builder.ticketNumber = ticketNumber
	return builder
}

func (builder *TicketDataBuilder) WithExitDate(exitDate string) *TicketDataBuilder {
	exitDateParse, _ := time.Parse(time.RFC3339, exitDate)
	builder.exitDate = exitDateParse
	return builder
}
func (builder *TicketDataBuilder) WithTotalTime(totalTime string) *TicketDataBuilder {
	builder.totalTime = totalTime
	return builder
}
func (builder *TicketDataBuilder) WithPaymentTotal(paymentTotal int64) *TicketDataBuilder {
	builder.paymentTotal = paymentTotal
	return builder
}

func (builder *TicketDataBuilder) Build() model.Ticket {
	return model.Ticket{
		TicketId:     builder.ticketId,
		TicketNumber: builder.ticketNumber,
		EnterDate:    builder.enterDate,
		ExitDate:     builder.exitDate,
		TotalTime:    builder.totalTime,
		PaymentTotal: builder.paymentTotal,
		SerialNumber: builder.serialNumber,
	}
}
