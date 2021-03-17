package infrastructure

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    controllers "money-send-api/interfaces/api"
)

func Init() {
    router := NewRouter()

    router.Use(middleware.Logger())
    router.Use(middleware.Recover())
    router.Logger.Fatal(router.Start(":1323"))
}

func NewRouter() *echo.Echo {
    e := echo.New()
    userController := controllers.NewUserController(NewSqlHandler())
    e.GET("/healthcheck", healthcheckHandler)
    e.POST("/users/add", userController.AddUser)
    e.GET("/users/list", userController.GetAllUsers)
    e.GET("/users/list/:user_id", userController.GetUser)
    e.PUT("/users/balance", userController.UpdateAllBalance)
    e.PUT("/users/balance/:user_id", userController.UpdateBalance)
    e.DELETE("/users/delete", userController.DeleteAllUser)
    e.DELETE("/users/delete/:user_id", userController.DeleteUser)

    return e
}

func healthcheckHandler(c echo.Context) error {
    return c.String(http.StatusOK, "healthcheck ok")
}
