package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type BikeAccessServiceMock struct {
	mock.Mock
}

func (mock *BikeAccessServiceMock) BikeAccess(bike model.Bike) error {
	args := mock.Called(bike)
	return args.Error(0)
}
