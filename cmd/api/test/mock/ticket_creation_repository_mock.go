package mock

import (
	"github.com/stretchr/testify/mock"
)

type TicketCreationRepositoryMock struct {
	mock.Mock
}

func (mock *TicketCreationRepositoryMock) SaveTicket(serialNumber string, enterDate string) (err error) {
	args := mock.Called(serialNumber, enterDate)
	return args.Error(0)
}
