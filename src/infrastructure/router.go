package infrastructure

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    controllers "money-send-api/interfaces/api"
    es "github.com/swaggo/echo-swagger"

    _ "money-send-api/docs"
)

func Init() {
    router := NewRouter()
    router.Use(middleware.Logger())
    router.Use(middleware.Recover())
    router.Logger.Fatal(router.Start(":1323"))
}

func NewRouter() *echo.Echo {
    e := echo.New()

    e.GET("/healthcheck", healthcheckHandler)

    userController := controllers.NewUserController(NewSqlHandler())

    e.POST("/login", userController.Login)
    e.POST("/regist", userController.Regist)

    api := e.Group("/api")
    api.Use(middleware.JWT([]byte("secret")))
    api.GET("/users/list", userController.GetAllUsers)
    api.GET("/users/list/:user_id", userController.GetUser)
    api.PUT("/users/balance", userController.UpdateAllBalance)
    api.PUT("/users/balance/:user_id", userController.UpdateBalance)
    api.DELETE("/users/delete/:user_id", userController.DeleteUser)

    e.GET("/swagger/*", es.WrapHandler)

    return e
}

func healthcheckHandler(c echo.Context) error {
    return c.String(http.StatusOK, "healthcheck ok")
}
