package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "github.com/labstack/echo"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHealthcheckHandler(t *testing.T) {
	e := echo.New()
	e.GET("/healthcheck", healthcheckHandler)

    req := httptest.NewRequest("GET", "/healthcheck", nil)
    rec := httptest.NewRecorder()

    e.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "healthcheck ok", rec.Body.String())
}
