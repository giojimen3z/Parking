package application_test

import (
	"errors"
	"testing"

	"github.com/Parking/cmd/api/app/application"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/Parking/cmd/api/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	bikeAccessServiceMock     = new(mock.BikeAccessServiceMock)
	ticketCreationServiceMock = new(mock.TicketCreationServiceMock)
)

func TestWhenAllBeOKAccessingBikeThenReturnNilError(t *testing.T) {
	bike := builder.NewBikeDataBuilder().Build()
	bikeAccessServiceMock.On("BikeAccess", bike).Return(nil).Once()
	ticketCreationServiceMock.On("TicketCreation", bike.SerialNumber).Return(nil).Once()
	bikeAccess := application.BikeAccess{
		BikeAccessService:     bikeAccessServiceMock,
		TicketCreationService: ticketCreationServiceMock,
	}

	err := bikeAccess.Handler(bike)

	assert.Nil(t, err)
	bikeAccessServiceMock.AssertExpectations(t)
	ticketCreationServiceMock.AssertExpectations(t)
}
func TestWhenFailedCreatingBikeThenReturnError(t *testing.T) {
	bike := builder.NewBikeDataBuilder().Build()

	expectedErrorMessage := errors.New("error getting repository information")
	bikeAccessServiceMock.On("BikeAccess", bike).Return(expectedErrorMessage).Once()
	bikeAccess := application.BikeAccess{
		BikeAccessService: bikeAccessServiceMock}

	err := bikeAccess.Handler(bike)

	assert.EqualError(t, err, expectedErrorMessage.Error())
	bikeAccessServiceMock.AssertExpectations(t)
}
func TestWhenFailedCreatingTicketThenReturnError(t *testing.T) {
	bike := builder.NewBikeDataBuilder().Build()
	bikeAccessServiceMock.On("BikeAccess", bike).Return(nil).Once()
	expectedErrorMessage := errors.New("error creating ticket by Bike with serial number :  MR145987D12")
	ticketCreationServiceMock.On("TicketCreation", bike.SerialNumber).Return(expectedErrorMessage).Once()
	bikeAccess := application.BikeAccess{
		BikeAccessService:     bikeAccessServiceMock,
		TicketCreationService: ticketCreationServiceMock,
	}

	err := bikeAccess.Handler(bike)

	assert.NotNil(t, err)
	bikeAccessServiceMock.AssertExpectations(t)
	ticketCreationServiceMock.AssertExpectations(t)
}
