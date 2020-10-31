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
	bikeAccessRepository = new(mock.BikeAccessRepositoryMock)
)

func TestWhenSendTheBikeIntoDToRepositoryThenShouldReturnOk(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	bikeAccessRepository.On("SaveBike", bike).Times(1).Return(nil)
	bikeAccessService := service.BikeAccessService{
		BikeAccessRepository: bikeAccessRepository,
	}
	err := bikeAccessService.BikeAccess(bike)

	assert.Nil(t, err)
	bikeAccessRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheBikeIntoDToRepositoryThenShouldReturnError(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	errorExpected := errors.New("error getting repository information")
	bikeAccessRepository.On("SaveBike", bike).Times(1).Return(errorExpected)
	bikeAccessService := service.BikeAccessService{
		BikeAccessRepository: bikeAccessRepository,
	}
	err := bikeAccessService.BikeAccess(bike)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	bikeAccessRepository.AssertExpectations(t)
}
