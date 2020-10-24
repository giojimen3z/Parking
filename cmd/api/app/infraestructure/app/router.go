package app

import (
	"os"

	"github.com/Parking/cmd/api/app/infraestructure/controller/middleware"
	"github.com/Parking/errorApi/logger"
	"github.com/Parking/errorApi/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()
	router.Use(middleware.ErrorHandler())

	mapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}
