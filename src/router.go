package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
    e := echo.New()

	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/regist", registUser)

	api := e.Group("/api")
    api.Use(middleware.JWTWithConfig(handler.Config))

    return e
}
