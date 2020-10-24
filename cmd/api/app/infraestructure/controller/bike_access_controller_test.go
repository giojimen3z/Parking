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

func setupReprocessClaimInsuranceController(bikeAccessMock *bikeMock.BikeAccessMock) (*gin.Engine, *controller.BikeAccessController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.BikeAccessController{BikeAccessApplication: bikeAccessMock}
}

func TestWhenMakeClaimReprocessThenReturn200(t *testing.T) {

	router, controllerBike := setupReprocessClaimInsuranceController(&bikeAccessMock)
	bikeAccessMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	claimsRouterGroup := router.Group(routerGroupBase)
	claimsRouterGroup.POST("/Bike/Access", controllerBike.MakeAccessBike)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, bikeAccessURITest, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	bikeAccessMock.AssertExpectations(t)
}

func TestWhenMakeClaimReprocessFailedThenReturn404(t *testing.T) {

	router, controllerBike := setupReprocessClaimInsuranceController(&bikeAccessMock)
	errorExpected := exception.DataNotFound{ErrMessage: "we didn't´t found information"}
	bikeAccessMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	claimsRouterGroup := router.Group(routerGroupBase)
	claimsRouterGroup.POST("/Bike/Access", controllerBike.MakeAccessBike)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, bikeAccessURITest, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	bikeAccessMock.AssertExpectations(t)
}
