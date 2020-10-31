package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingCreationMock struct {
	mock.Mock
}

func (mock *ParkingCreationMock) Handler(parking model.Parking) (err error) {
	args := mock.Called(parking)
	return args.Error(0)
}
