package infrastructure

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    controllers "money-send-api/interfaces/api"
)

func Init() {
    // Echo initialize
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Create user controller
    userController := controllers.NewUserController(NewSqlHandler())

    // Add routings
    e.GET("/healthcheck", healthcheckHandler)

    e.POST("/users/add", userController.AddUser)

    e.GET("/users/list", userController.GetAllUsers)
    e.GET("/users/list/:user_id", userController.GetUser)

    e.PUT("/users/balance", userController.UpdateAllBalance)
    e.PUT("/users/balance/:user_id", userController.UpdateBalance)

    e.DELETE("/users/delete/:user_id", userController.DeleteUser)
 
    e.Logger.Fatal(e.Start(":1323"))
}

func healthcheckHandler(c echo.Context) error {
    return c.String(http.StatusOK, "healthcheck ok")
}
