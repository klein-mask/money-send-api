package controllers

import (
    "net/http"
    "strconv"
    "money-send-api/domain"
    "money-send-api/interfaces/database"
    "money-send-api/usecase"

    "github.com/labstack/echo"
)

type UserController struct {
    Interactor usecase.UserInteractor
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
    balance, _ := strconv.ParseInt(c.FormValue("balance"), 10, 64)
    err := controller.Interactor.UpdateAllBalance(balance)
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