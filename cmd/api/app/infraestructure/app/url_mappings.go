package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_insurance-backoffice-backend/cmd/api/app/infrastructure/controller"
)

func mapUrls(router *gin.Engine) {

	router.GET("/ping", controller.PingController.Ping)
}