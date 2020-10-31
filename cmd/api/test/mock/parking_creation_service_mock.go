package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingCreationServiceMock struct {
	mock.Mock
}

func (mock *ParkingCreationServiceMock) ParkingCreation(parking model.Parking) error {
	args := mock.Called(parking)
	return args.Error(0)
}
