package app

import (
	"fmt"

	"github.com/Parking/cmd/api/app/infraestructure/config"
	"github.com/Parking/cmd/api/app/infraestructure/container"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping", controller.PingController.Ping)
	prefix := fmt.Sprintf("%s/api/Parking", prefixScope)

	baseUrl := router.Group(prefix)
	bike := baseUrl.Group("/Bike")
	bike.POST("/Access", container.GetContentController().MakeAccessBike)
}
