package controller

import (
	"fmt"
	"net/http"

	"github.com/Parking/cmd/api/app/application"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/gin-gonic/gin"
)


const (
	ParkingCreatedMsg = "the Parking  %v  was created successfully"
)

// ParkingCreationController  used for inject the use case
type ParkingCreationController struct {
	ParkingCreationApplication application.ParkingCreationApplication
}

//MakeParkingCreation is to execute the use case for create the parking
func (parkingCreationController *ParkingCreationController) MakeParkingCreation(context *gin.Context) {

	parking := parkingCreationController.mapParking(context)

	err := parkingCreationController.ParkingCreationApplication.Handler(parking)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, fmt.Sprintf(BikeAccessGranted, parking.ParkingName))

}

// mapParking used for get attributes for  any parking from url
func (parkingCreationController *ParkingCreationController) mapParking(context *gin.Context) (parking model.Parking) {

	parking = model.Parking{
		ParkingName:    context.Query("parking_name"),
		ParkingAddress: context.Query("parking_address"),
		ParkingOwner:   context.Query("parking_owner"),
	}
	return parking
}
