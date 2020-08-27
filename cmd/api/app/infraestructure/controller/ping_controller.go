package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/_integrations/nrgin/v1"
)

const (
	pong string = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(c *gin.Context)
}

type pingController struct{}

func (controller *pingController) Ping(c *gin.Context) {
	if txn := nrgin.Transaction(c); txn != nil {
		_ = txn.Ignore()
	}

	c.String(http.StatusOK, pong)
}