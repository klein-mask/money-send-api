package infrastructure

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    controllers "money-send-api/src/interfaces/api"
)

func Init() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    userController := controllers.NewUserController(NewSqlHandler())

    e.GET("/ping", func(c echo.Context) error {
        return c.String(http.StatusOK, "ping ok")
    })

    e.GET("/users", func(c echo.Context) error {
        users := userController.GetUser() 
        c.Bind(&users) 
        return c.JSON(http.StatusOK, users)
    })

    e.POST("/users", func(c echo.Context) error {
        userController.Create(c)
        return c.String(http.StatusOK, "created")
    })
    /*

    e.GET("/users", func(c echo.Context) error {
        users := userController.GetUser() 
        c.Bind(&users) 
        return c.JSON(http.StatusOK, users)
    })



    e.DELETE("/users/:id", func(c echo.Context) error {
        id := c.Param("id")
        userController.Delete(id)
        return c.String(http.StatusOK, "deleted")
    })
    */
    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}
