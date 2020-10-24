
package mlhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Parking/errorApi/apierrors"
	"github.com/gin-gonic/gin"
)

var production bool = os.Getenv("GO_ENVIRONMENT") == "production"

func DefaultRouter() *gin.Engine {
	return CustomRouter(MeliRouterConfig{NewRelicHttpStatusToIgnore: []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusMethodNotAllowed}})
}

func CustomRouter(conf MeliRouterConfig) *gin.Engine {
	router := gin.New()

	if conf.DisableCancellationOnClientDisconnect {
		router.Use(func(c *gin.Context) {
			c.Request = c.Request.WithContext(context.Background())
			c.Next()
		})
	}
	if !production {
		router.Use(gin.Logger())
	}
	router.NoRoute(noRouteHandler)
	return router
}

type MeliRouterConfig struct {
	DisableDataDogMetrics            bool
	DisableNewRelicMetrics           bool
	DisablePanicRecover              bool
	DisableCommonApiFilter           bool
	DisablePprof                     bool
	EnableResponseCompressionSupport bool
	DisableHeaderForwarding          bool
	NewRelicHttpStatusToIgnore       []int

	// DisableCancellationOnClientDisconnect tells the server to detach the
	// c.Request.Context() from the incoming TCP connection. If set to false
	// then the client closing the connection does not cancels the context.
	// The default behavior from Go is to cancel the request context if it can
	// ensure that there's no one on the other side to read the response.
	DisableCancellationOnClientDisconnect bool
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, apierrors.NewNotFoundApiError(fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)))
}

func AddResponseExpiration(time time.Duration, c *gin.Context) {
	var roundTime int = int(time.Seconds())
	c.Writer.Header()["Cache-Control"] = []string{fmt.Sprintf("max-age=%v,stale-while-revalidate=%v, stale-if-error=%v", roundTime, roundTime/2, roundTime*2)}
}