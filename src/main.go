package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type User struct {
    Id string
    Name string
    Password
}

func main() {
    router := newRouter()
    router.Logger.Fatal(router.Start(":8080"))
}



func helloHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Hello")
}