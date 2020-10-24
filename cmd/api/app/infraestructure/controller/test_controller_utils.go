package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func GetMockedContext(method string, url string, requestBody io.Reader, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(method, url, requestBody)
	return c
}

func GetTargetResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func RunRequestWithHeaders(t *testing.T, router *gin.Engine, httpMethod string, url string, payload string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	request, err := http.NewRequest(httpMethod, url, bytes.NewReader([]byte(payload)))
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	if err != nil {
		panic(fmt.Sprintf("An error occurred while creating the request for URL (%s)", url))
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, request)
	return res
}
