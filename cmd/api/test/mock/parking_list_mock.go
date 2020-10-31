package mock

import (
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type ParkingListMock struct {
	mock.Mock
}

func (mock *ParkingListMock) Handler() (parkingList []model.Parking, err error) {
	args := mock.Called()
	return args.Get(0).([]model.Parking), args.Error(1)
}
