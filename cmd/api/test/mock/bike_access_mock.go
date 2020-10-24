package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type BikeAccessMock struct {
	mock.Mock
}

func (mock *BikeAccessMock) Handler(bike model.Bike) (err error) {
	args := mock.Called(bike)
	return args.Error(0)
}
