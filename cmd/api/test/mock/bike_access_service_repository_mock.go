package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type BikeAccessRepositoryMock struct {
	mock.Mock
}

func (mock *BikeAccessRepositoryMock)SaveBike(bike model.Bike) (err error) {
	args := mock.Called(bike)
	return args.Error(0)
}
