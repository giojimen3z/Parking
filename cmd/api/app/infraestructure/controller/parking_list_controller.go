package controller

import (
	"net/http"

	"github.com/Parking/cmd/api/app/application"
	"github.com/Parking/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

const (
	failedListParking = "we didn't  list all parking"
)

// ParkingListController  used for inject the use case
type ParkingListController struct {
	ParkingListApplication application.ParkingListApplication
}

//MakeParkingList is to execute the use case for get all parking
func (parkingListController *ParkingListController) MakeParkingList(context *gin.Context) {

	parkingList, err := parkingListController.ParkingListApplication.Handler()

	if err != nil {
		err := apierrors.NewNotFoundApiError(failedListParking)
		context.JSON(err.Status(), err)
	}

	context.JSON(http.StatusOK, parkingList)

}
