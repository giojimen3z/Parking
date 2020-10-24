package container

import (
	"github.com/Parking/cmd/api/app/domain/service"
)

func getBikeAccessService() service.BikeAccessServicePort {
	return &service.BikeAccessService{
		BikeAccessRepository: getBikeAccessRepository(),
	}
}
