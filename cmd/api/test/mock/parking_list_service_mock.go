package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingListServiceMock struct {
	mock.Mock
}

func (mock *ParkingListServiceMock) ParkingList() (parkingList []model.Parking, err error) {
	args := mock.Called()
	return args.Get(0).([]model.Parking), args.Error(1)
}
