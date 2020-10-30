package mock

import (
	"github.com/stretchr/testify/mock"
)

type TicketCreationServiceMock struct {
	mock.Mock
}

func (mock *TicketCreationServiceMock) TicketCreation(serialNumber string) error {
	args := mock.Called(serialNumber)
	return args.Error(0)
}
