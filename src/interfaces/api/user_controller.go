package controllers

import (
    "money-send-api/src/domain"
    "money-send-api/src/interfaces/database"
    "money-send-api/src/usecase"

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

func (controller *UserController) Create(c echo.Context) {
    u := domain.User{}
    c.Bind(&u)
    controller.Interactor.Add(u)
    createdUsers := controller.Interactor.GetInfo()
    c.JSON(201, createdUsers)
    return
}

func (controller *UserController) GetUsers() []domain.User {
    res := controller.Interactor.GetInfo()
    return res
}

func (controller *UserController) SelectUser(id string) domain.User {
    res := controller.Interactor.SelectUser(id)
    return res
}

func (controller *UserController) UpdateBalance(id string, balance int64) {
    controller.Interactor.UpdateBalance(id, balance)
}

func (controller *UserController) Delete(id string) {
    controller.Interactor.Delete(id)
}