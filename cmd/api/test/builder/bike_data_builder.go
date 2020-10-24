package builder

import "github.com/Parking/cmd/api/app/domain/model"

type BikeDataBuilder struct {
	bikeId       int64
	serialNumber string
	brand        string
	color        string
}

func NewBikeDataBuilder() *BikeDataBuilder {
	return &BikeDataBuilder{
		bikeId:       1,
		serialNumber: "MR145987D12",
		brand:        "GW",
		color:        "Negro",
	}
}
func (builder *BikeDataBuilder) Build() model.Bike {
	return model.Bike{
		BikeId:       builder.bikeId,
		SerialNumber: builder.serialNumber,
		Brand:        builder.brand,
		Color:        builder.color,
	}
}
