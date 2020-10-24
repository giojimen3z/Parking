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
	bikeAccessServiceMock = new(mock.BikeAccessServiceMock)
)

func TestWhenAllBeOKThenReturnNilError(t *testing.T) {

	bike := builder.NewBikeDataBuilder().Build()
	bikeAccessServiceMock.On("BikeAccess", bike).Return(nil).Once()

	bikeAccess := application.BikeAccess{
		BikeAccessService: bikeAccessServiceMock	}

	err := bikeAccess.Handler(bike)

	assert.Nil(t, err)
	bikeAccessServiceMock.AssertExpectations(t)
}
func TestWhenRepositoryFailedThenReturnError(t *testing.T) {
	bike := builder.NewBikeDataBuilder().Build()

	expectedErrorMessage := errors.New("error getting repository information")
	bikeAccessServiceMock.On("BikeAccess", bike).Return(expectedErrorMessage).Once()
	bikeAccess := application.BikeAccess{
		BikeAccessService: bikeAccessServiceMock	}

	err := bikeAccess.Handler(bike)

	assert.EqualError(t, err, expectedErrorMessage.Error())
	bikeAccessServiceMock.AssertExpectations(t)
}

