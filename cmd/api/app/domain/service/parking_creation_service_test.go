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
	parkingCreationRepository = new(mock.ParkingCreationRepositoryMock)
)

func TestWhenSendTheParkingToRepositoryThenShouldReturnOk(t *testing.T) {

	parking := builder.NewParkingDataBuilder().Build()
	parkingCreationRepository.On("SaveParking", parking).Times(1).Return(nil)
	parkingCreationService := service.ParkingCreationService{
		ParkingCreationRepository: parkingCreationRepository,
	}
	err := parkingCreationService.ParkingCreation(parking)

	assert.Nil(t, err)
	parkingCreationRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheParkingToRepositoryThenShouldReturnError(t *testing.T) {

	parking := builder.NewParkingDataBuilder().Build()
	errorExpected := errors.New("error getting repository information")
	parkingCreationRepository.On("SaveParking", parking).Times(1).Return(errorExpected)
	parkingCreationService := service.ParkingCreationService{
		ParkingCreationRepository: parkingCreationRepository,
	}
	err := parkingCreationService.ParkingCreation(parking)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	parkingCreationRepository.AssertExpectations(t)
}
