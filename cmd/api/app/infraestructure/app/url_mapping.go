package app

import (
	"fmt"

	"github.com/Parking/cmd/api/app/infraestructure/config"
	"github.com/Parking/cmd/api/app/infraestructure/container"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping", controller.PingController.Ping)
	prefix := fmt.Sprintf("%s/api/Parking", prefixScope)

	baseUrl := router.Group(prefix)
	parking := baseUrl.Group("")
	bike := baseUrl.Group("/Bike")
	bike.POST("/Access", container.GetBikeAccessController().MakeAccessBike)
	parking.POST("", container.GetParkingCreationController().MakeParkingCreation)
	parking.GET("/", container.GetParkingListController().MakeParkingList)
}
