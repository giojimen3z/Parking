package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type TicketCreationRepositoryMock struct {
	mock.Mock
}

func (mock *TicketCreationRepositoryMock) SaveTicket(ticket model.Ticket) (err error) {
	args := mock.Called(ticket)
	return args.Error(0)
}
