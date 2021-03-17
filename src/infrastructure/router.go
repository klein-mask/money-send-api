package infrastructure

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "strconv"
    controllers "money-send-api/src/interfaces/api"
)

func Init() {
    // Echo initialize
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Create user controller
    userController := controllers.NewUserController(NewSqlHandler())

    // Add routings
    e.GET("/healthcheck", func(c echo.Context) error {
        return c.String(http.StatusOK, "healthcheck ok")
    })

    e.POST("/users", func(c echo.Context) error {
        userController.AddUser(c)
        return c.String(http.StatusOK, "created")
    })

    e.GET("/users", func(c echo.Context) error {
        users := userController.GetAllUsers() 
        c.Bind(&users) 
        return c.JSON(http.StatusOK, users)
    })

    e.GET("/users/:user_id", func(c echo.Context) error {
        user_id := c.Param("user_id")
        user := userController.GetUser(user_id) 
        c.Bind(&user) 
        return c.JSON(http.StatusOK, user)
    })

    e.PUT("/users/balance/all", func(c echo.Context) error {
        barance := c.FormValue("balance")
        balance64, _ := strconv.ParseInt(barance, 10, 64)
        userController.UpdateAllBalance(balance64)
        return c.String(http.StatusOK, "updated all user balance")
    })

    e.PUT("/users/balance/:user_id", func(c echo.Context) error {
        user_id := c.Param("user_id")
        barance := c.FormValue("balance")
        balance64, _ := strconv.ParseInt(barance, 10, 64)
        userController.UpdateBalance(user_id, balance64)
        return c.String(http.StatusOK, "updated user balance")
    })

    e.DELETE("/users/delete/:id", func(c echo.Context) error {
        id := c.Param("id")
        userController.DeleteUser(id)
        return c.String(http.StatusOK, "deleted")
    })
 
    e.Logger.Fatal(e.Start(":1323"))
}
