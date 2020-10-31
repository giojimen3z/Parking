package builder

import "github.com/Parking/cmd/api/app/domain/model"

type ParkingDataBuilder struct {
	parkingId      int64
	parkingName    string
	parkingAddress string
	parkingOwner   string
}

func NewParkingDataBuilder() *ParkingDataBuilder {
	return &ParkingDataBuilder{
		parkingId:      1,
		parkingName:    "PArkAutosBosa",
		parkingAddress: "Calle 63 sur # 45-24",
		parkingOwner:   "Andres Rodriguez",
	}
}
func (builder *ParkingDataBuilder) Build() model.Parking {
	return model.Parking{
		ParkingId:      builder.parkingId,
		ParkingName:    builder.parkingName,
		ParkingAddress: builder.parkingAddress,
		ParkingOwner:   builder.parkingOwner,
	}
}
