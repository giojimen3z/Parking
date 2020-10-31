package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/Parking/cmd/api/app/infraestructure/controller/middleware"
	bikeMock "github.com/Parking/cmd/api/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	bikeAccessURITest = "/api/Parking/Bike/Access"
	routerGroupBase   = "/api/Parking"
)

var (
	bikeAccessMock bikeMock.BikeAccessMock
)

func setupBikeAccessController(bikeAccessMock *bikeMock.BikeAccessMock) (*gin.Engine, *controller.BikeAccessController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.BikeAccessController{BikeAccessApplication: bikeAccessMock}
}

func TestWhenMakeBikeAccessThenReturn200(t *testing.T) {

	router, controllerBike := setupBikeAccessController(&bikeAccessMock)
	bikeAccessMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	bikeRouterGroup := router.Group(routerGroupBase)
	bikeRouterGroup.POST("/Bike/Access", controllerBike.MakeAccessBike)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, bikeAccessURITest, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	bikeAccessMock.AssertExpectations(t)
}

func TestWhenMakeBikeAccessFailedThenReturn404(t *testing.T) {

	router, controllerBike := setupBikeAccessController(&bikeAccessMock)
	errorExpected := exception.DataNotFound{ErrMessage: "we didn't found information"}
	bikeAccessMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	bikeRouterGroup := router.Group(routerGroupBase)
	bikeRouterGroup.POST("/Bike/Access", controllerBike.MakeAccessBike)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, bikeAccessURITest, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	bikeAccessMock.AssertExpectations(t)
}
