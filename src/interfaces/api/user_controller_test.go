package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func TestAddUser(t *testing.T) {
    e := echo.New()
    req := new(http.Request)
    rec := httptest.NewRecorder()
    c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
    c.SetPath("/users/add")
    h := Handler{testMockDB}

    h.GetUser(c)
    if rec.Body.String() != "{\"id\":\"one\"}" {
        t.Errorf("expected response Id: one, got %s", rec.Body.String())
    }
}
