package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/Parking/cmd/api/app/infraestructure/controller/middleware"
	parkingMock "github.com/Parking/cmd/api/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	parkingCreationURITest = "/api/Parking"
)

var (
	parkingCreationMock parkingMock.ParkingCreationMock
)

func setupParkingCreationController(parkingCreationMock *parkingMock.ParkingCreationMock) (*gin.Engine, *controller.ParkingCreationController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.ParkingCreationController{ParkingCreationApplication: parkingCreationMock}
}
func TestWhenMakeParkingCreationThenReturn200(t *testing.T) {

	router, controllerParking := setupParkingCreationController(&parkingCreationMock)
	parkingCreationMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	parkingRouterGroup := router.Group(parkingCreationURITest)
	parkingRouterGroup.POST("", controllerParking.MakeParkingCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, parkingCreationURITest, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	bikeAccessMock.AssertExpectations(t)
}

func TestWhenMakeParkingCreationFailedThenReturn404(t *testing.T) {

	router, controllerParking := setupParkingCreationController(&parkingCreationMock)
	errorExpected := exception.DataNotFound{ErrMessage: "we didn't found information"}
	parkingCreationMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	parkingRouterGroup := router.Group(parkingCreationURITest)
	parkingRouterGroup.POST("", controllerParking.MakeParkingCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, parkingCreationURITest, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	bikeAccessMock.AssertExpectations(t)
}
