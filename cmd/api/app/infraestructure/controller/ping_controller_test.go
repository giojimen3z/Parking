package controller_test

import (
	"net/http"
	"testing"

	"github.com/Parking/cmd/api/app/infraestructure/controller"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "pong", controller.Pong)
}

func TestPing(t *testing.T) {
	response := controller.GetTargetResponse()
	c := controller.GetMockedContext(http.MethodGet, "/ping", nil, response)

	controller.PingController.Ping(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "pong", response.Body.String())
}
