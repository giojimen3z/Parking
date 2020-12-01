package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/cmd/api/app/domain/model"
	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/Parking/cmd/api/app/infraestructure/controller/middleware"
	"github.com/Parking/cmd/api/test/builder"
	parkingMock "github.com/Parking/cmd/api/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	parkingListURITest = "/api/Parking/"
)

var (
	parkingListMock parkingMock.ParkingListMock
)

func setupParkingListController(parkingListMock *parkingMock.ParkingListMock) (*gin.Engine, *controller.ParkingListController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.ParkingListController{ParkingListApplication: parkingListMock}
}
func TestWhenMakeParkingListThenReturn200AndParkingList(t *testing.T) {
	router, controllerParkingList := setupParkingListController(&parkingListMock)
	parkingList := []model.Parking{builder.NewParkingDataBuilder().Build()}
	expectedBody, err := json.Marshal(parkingList)
	parkingListMock.On("Handler").Times(1).Return(parkingList, nil).Once()
	parkingListRouterGroup := router.Group(routerGroupBase)
	parkingListRouterGroup.GET("/", controllerParkingList.MakeParkingList)

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, parkingListURITest, "", nil)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(expectedBody), response.Body.String())
	bikeAccessMock.AssertExpectations(t)
}

func TestWhenMakeParkingListThenReturn404(t *testing.T) {

	router, controllerParkingList := setupParkingListController(&parkingListMock)
	parkingList := []model.Parking{}
	errorExpected := exception.DataNotFound{ErrMessage: "{\"message\":\"we didn't  list all parking\",\"error\":\"not_found\",\"status\":404,\"cause\":[]}[]"}
	parkingListMock.On("Handler").Times(1).Return(parkingList, errorExpected).Once()
	parkingListRouterGroup := router.Group(routerGroupBase)
	parkingListRouterGroup.GET("/", controllerParkingList.MakeParkingList)

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, parkingListURITest, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, errorExpected.Error(), response.Body.String())
	parkingListMock.AssertExpectations(t)
}
