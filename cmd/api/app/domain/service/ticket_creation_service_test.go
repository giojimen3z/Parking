package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/Parking/cmd/api/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	ticketCreationRepository = new(mock.TicketCreationRepositoryMock)
)

func TestWhenSaveTheTicketIntoDBAThenShouldReturnOk(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	enterDate := time.Now().UTC().Format(time.RFC3339)
	ticketCreationRepository.On("SaveTicket", bike.SerialNumber,enterDate).Times(1).Return(nil)
	ticketCreationService := service.TicketCreationService{
		TicketCreationRepository : ticketCreationRepository,
	}
	err := ticketCreationService.TicketCreation(bike.SerialNumber)

	assert.Nil(t, err)
	ticketCreationRepository.AssertExpectations(t)
}
func TestWhenFailedSaveTheTicketIntoDBAThenShouldReturnError(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	errorExpected := errors.New("error getting repository information")
	enterDate := time.Now().UTC().Format(time.RFC3339)
	ticketCreationRepository.On("SaveTicket", bike.SerialNumber,enterDate).Times(1).Return(errorExpected)
	ticketCreationService := service.TicketCreationService{
		TicketCreationRepository : ticketCreationRepository,
	}
	err := ticketCreationService.TicketCreation(bike.SerialNumber)


	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	ticketCreationRepository.AssertExpectations(t)
}
