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
	parkingCreationServiceMock = new(mock.ParkingCreationServiceMock)
)

func TestWhenAllBeOKCreatingParkingThenReturnNilError(t *testing.T) {

	parking := builder.NewParkingDataBuilder().Build()
	parkingCreationServiceMock.On("ParkingCreation", parking).Return(nil).Once()
	parkingCreation := application.ParkingCreation{
		ParkingCreationService: parkingCreationServiceMock,
	}

	err := parkingCreation.Handler(parking)

	assert.Nil(t, err)
	parkingCreationServiceMock.AssertExpectations(t)
}
func TestWhenFailedCreatingParkingThenReturnError(t *testing.T) {
	parking := builder.NewParkingDataBuilder().Build()
	expectedErrorMessage := errors.New("error getting repository information")
	parkingCreationServiceMock.On("ParkingCreation", parking).Return(expectedErrorMessage).Once()
	parkingCreation := application.ParkingCreation{
		ParkingCreationService: parkingCreationServiceMock,
	}

	err := parkingCreation.Handler(parking)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	parkingCreationServiceMock.AssertExpectations(t)
}
