package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingListRepositoryMock struct {
	mock.Mock
}

func (mock *ParkingListRepositoryMock) ListParking() (parkingList []model.Parking, err error) {
	args := mock.Called()
	return args.Get(0).([]model.Parking), args.Error(1)
}
