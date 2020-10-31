package service_test

import (
	"errors"
	"testing"

	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/domain/service"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/Parking/cmd/api/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	parkingListRepository = new(mock.ParkingListRepositoryMock)
)

func TestWhenGetTheParkingListFromRepositoryThenShouldReturnOk(t *testing.T) {
	parkingLots := []model.Parking{builder.NewParkingDataBuilder().Build()}
	parkingListRepository.On("ListParking").Times(1).Return(parkingLots, nil)
	parkingListService := service.ParkingListService{
		ParkingListRepository: parkingListRepository,
	}

	parking, err := parkingListService.ParkingList()

	assert.Nil(t, err)
	assert.Equal(t, parkingLots, parking)
	parkingListRepository.AssertExpectations(t)
}
func TestWhenFailedGetTheParkingListFromRepositoryThenShouldReturnError(t *testing.T) {

	parkingLots := []model.Parking{}
	errorExpected := errors.New("Error Getting the parking list from repository")
	parkingListRepository.On("ListParking").Times(1).Return(parkingLots, errorExpected)
	parkingListService := service.ParkingListService{
		ParkingListRepository: parkingListRepository,
	}

	parking, err := parkingListService.ParkingList()

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	assert.Equal(t, parkingLots, parking)
	parkingListRepository.AssertExpectations(t)
}
