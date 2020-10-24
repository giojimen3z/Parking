package controller

import (
	"fmt"
	"net/http"

	"github.com/Parking/cmd/api/app/application"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/gin-gonic/gin"
)

//URIBikeAccess is for the endpoint to add new bike
const (
	BikeAccessGranted = "the bike with the serial number %v was access successfully"
)

// BikeAccessController  used for inject the use case
type BikeAccessController struct {
	BikeAccessApplication application.AccessBikeApplication
}


//MakeSyncTranslation is to execute the use case for sync translation
func (bikeAccessController *BikeAccessController) MakeAccessBike(context *gin.Context) {

	bike := bikeAccessController.mapBike(context)

	err := bikeAccessController.BikeAccessApplication.Handler(bike)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, fmt.Sprintf(BikeAccessGranted, bike.SerialNumber))

}

// mapBike used for get attributes for  any bike from url
func (bikeAccessController *BikeAccessController) mapBike(context *gin.Context) (bike model.Bike) {

	bike = model.Bike{
		SerialNumber: context.Query("serial_number"),
		Brand:        context.Query("brand"),
		Color:        context.Query("color"),
	}
	return bike
}


func abort(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
