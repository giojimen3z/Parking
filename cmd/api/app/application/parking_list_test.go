package application_test

import (
	"errors"
	"testing"

	"github.com/Parking/cmd/api/app/application"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/test/builder"
	"github.com/Parking/cmd/api/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	parkingListServiceMock = new(mock.ParkingListServiceMock)
)

func TestWhenAllBeOKAndListParkingThenReturnNilError(t *testing.T) {

	parkingLots := []model.Parking{builder.NewParkingDataBuilder().Build()}
	parkingListServiceMock.On("ParkingList").Return(parkingLots,nil).Once()
	parkingList := application.ParkingList{
		ParkingListService: parkingListServiceMock,
	}

	parking,err := parkingList.Handler()

	assert.Nil(t, err)
	assert.Equal(t,parkingLots,parking)
	parkingListServiceMock.AssertExpectations(t)
}
func TestWhenFailedListParkingThenReturnError(t *testing.T) {
	parkingLots := []model.Parking{}
	expectedErrorMessage := errors.New("error getting information from service Parking List")
	parkingListServiceMock.On("ParkingList").Return(parkingLots,expectedErrorMessage).Once()
	parkingList := application.ParkingList{
		ParkingListService: parkingListServiceMock,
	}

	parking,err := parkingList.Handler()

	assert.NotNil(t,err)
	assert.Equal(t,parkingLots,parking)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	parkingCreationServiceMock.AssertExpectations(t)
}
