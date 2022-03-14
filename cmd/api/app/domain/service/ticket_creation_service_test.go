package service_test

import (
	"errors"
	"testing"

	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/Parking/cmd/api/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	ticketCreationRepository = new(mock.TicketCreationRepositoryMock)
)

func TestWhenFailedSendTheTicketToRepositoryThenShouldReturnError(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	ticket := builder.NewTicketDataBuilder().WithTickedID(0).WithTicketNumber("").WithExitDate("0001-01-01 00:00:00 +0000").WithPaymentTotal(0).WithTotalTime("").Build()
	errorExpected := errors.New("error getting repository information")
	ticketCreationRepository.On("SaveTicket", ticket).Times(1).Return(errorExpected)
	ticketCreationService := service.TicketCreationService{
		TicketCreationRepository: ticketCreationRepository,
	}

	err := ticketCreationService.TicketCreation(bike.SerialNumber)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	ticketCreationRepository.AssertExpectations(t)
}

func TestWhenSendTheTicketToRepositoryThenShouldReturnOk(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	ticket := builder.NewTicketDataBuilder().WithTickedID(0).WithTicketNumber("").WithExitDate("0001-01-01 00:00:00 +0000").WithPaymentTotal(0).WithTotalTime("").Build()
	ticketCreationRepository.On("SaveTicket", ticket).Times(1).Return(nil)
	ticketCreationService := service.TicketCreationService{
		TicketCreationRepository: ticketCreationRepository,
	}
	err := ticketCreationService.TicketCreation(bike.SerialNumber)

	assert.Nil(t, err)
	ticketCreationRepository.AssertExpectations(t)
}
