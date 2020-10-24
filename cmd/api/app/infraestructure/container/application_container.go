package container

import (
	"github.com/Parking/cmd/api/app/application"
)

func GetBikeAccessApplication() application.AccessBikeApplication {
	return &application.BikeAccess{  BikeAccessService : getBikeAccessService()}
}
