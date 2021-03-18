package controllers

import (
    "net/http"
    "strconv"
    "money-send-api/domain"
    "money-send-api/interfaces/database"
    "money-send-api/usecase"

    "github.com/labstack/echo"
    _ "fmt"
    _ "encoding/json"
)

type UserController struct {
    Interactor usecase.UserInteractor
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

func (controller *UserController) AddUser(c echo.Context) error {
    u := domain.User{}
    c.Bind(&u)

    err := controller.Interactor.AddUser(u)
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
    //bytes, _ := json.Marshal(jsonData) // jsonのバイナリが出力される
    //fmt.Println(string(bytes))
    err := controller.Interactor.UpdateAllBalance(jsonData.Balance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated all user's balance."
    return c.JSON(http.StatusOK, msg)
}


func (controller *UserController) UpdateBalance(c echo.Context) error {
    userId := c.Param("user_id")
    barance, _ := strconv.ParseInt(c.FormValue("balance"), 10, 64)
    err := controller.Interactor.UpdateBalance(userId, barance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated balance User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}

func (controller *UserController) DeleteUser(c echo.Context) error {
    userId := c.Param("user_id")
    err := controller.Interactor.DeleteUser(userId)
    if err != nil {
        return err
    }
    msg := "[Success] : Deleted User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}