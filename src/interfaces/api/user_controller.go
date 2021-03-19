package controllers

import (
    "net/http"
    "money-send-api/domain"
    "money-send-api/interfaces/database"
    "money-send-api/usecase"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "time"
    _ "fmt"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

type LoginUser struct {
    Name string `json:"name"`
    Password string `json:"password"`
}

type JsonData struct {
    Balance int64 `json:"balance"`
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Login(c echo.Context) error {
    lu := LoginUser{}
    c.Bind(&lu)

    name := lu.Name
    password := lu.Password

    err := controller.Interactor.Login(name, password)

    //fmt.Println(err)
    if err != nil {
        return err
    }
    token := jwt.New(jwt.SigningMethodHS256)

    // set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["name"] = name
    claims["admin"] = true
    claims["iat"] = time.Now().Unix()
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // generate encoded token and send it as response
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })

    return echo.ErrUnauthorized
}

func (controller *UserController) Regist(c echo.Context) error {
    u := domain.User{}
    c.Bind(&u)

    err := controller.Interactor.Regist(u)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, u)
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
    users, err := controller.Interactor.GetAllUsers()
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, users)
}

func (controller *UserController) GetUser(c echo.Context) error {
    userId := c.Param("user_id")
    user, err := controller.Interactor.GetUser(userId)
    if err != nil {
        return err
    } else if user.Name == "" {
        msg := "[Not Found] : User id <" + userId + "> is not found."
        return c.JSON(http.StatusOK, msg)
    }
    return c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateAllBalance(c echo.Context) error {
    jsonData := JsonData{}
    c.Bind(&jsonData)
    err := controller.Interactor.UpdateAllBalance(jsonData.Balance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated all user's balance."
    return c.JSON(http.StatusOK, msg)
}


func (controller *UserController) UpdateBalance(c echo.Context) error {
    userId := c.Param("user_id")
    jsonData := JsonData{}
    c.Bind(&jsonData)
    err := controller.Interactor.UpdateBalance(userId, jsonData.Balance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated balance User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}

func (controller *UserController) DeleteUser(c echo.Context) error {
    userId := c.Param("user_id")
    err := controller.Interactor.DeleteUser(userId)
    //fmt.Println(string(userId))
    if err != nil {
        return err
    }
    msg := "[Success] : Deleted User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}