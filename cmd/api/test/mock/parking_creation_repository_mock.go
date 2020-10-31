package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingCreationRepositoryMock struct {
	mock.Mock
}

func (mock *ParkingCreationRepositoryMock) SaveParking(parking model.Parking) (err error) {
	args := mock.Called(parking)
	return args.Error(0)
}
