package port

import "github.com/Parking/cmd/api/app/domain/model"

// ParkingCreationRepository interface to connect Parking Creation  implementation
type ParkingCreationRepository interface {
	// SaveParking post the Parking into DBA
	SaveParking(parking model.Parking) (err error)
}
