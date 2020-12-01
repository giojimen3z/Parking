package middleware

import (
	"net/http"

	"github.com/Parking/cmd/api/app/domain/exception"
	"github.com/Parking/pkg/apierrors"
	"github.com/Parking/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const internalServerErrorMessage = "an error occurred during the processing of your request"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		cause := errors.Cause(err.Err)

		if _, ok := cause.(exception.NotFound); ok {
			throwException(c, http.StatusNotFound, err.Err, cause)
			return
		}

		if _, ok := cause.(exception.InternalServerErrorPort); ok {
			throwException(c, http.StatusInternalServerError, err.Err, cause)
			return
		}

		logger.Error("middleware error 500", cause)
		throwException(c, http.StatusInternalServerError, errors.New(internalServerErrorMessage), cause)
	}
}

func throwException(ctx *gin.Context, status int, err error, cause error) {
	restErr := apierrors.NewApiError(err.Error(), http.StatusText(status), status, apierrors.CauseList{cause})
	logger.Error(restErr.Message(), cause)
	ctx.JSON(restErr.Status(), restErr)
}
