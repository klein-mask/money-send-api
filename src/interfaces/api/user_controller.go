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

func (controller *UserController) AddUser(c echo.Context) {
    u := domain.User{}
    c.Bind(&u)
    controller.Interactor.AddUser(u)
    users := controller.Interactor.GetAllUsers()
    c.JSON(201, users)
    return
}

func (controller *UserController) GetAllUsers() []domain.User {
    res := controller.Interactor.GetAllUsers()
    return res
}

func (controller *UserController) GetUser(id string) domain.User {
    res := controller.Interactor.GetUser(id)
    return res
}

func (controller *UserController) UpdateAllBalance(balance int64) {
    controller.Interactor.UpdateAllBalance(balance)
}


func (controller *UserController) UpdateBalance(id string, balance int64) {
    controller.Interactor.UpdateBalance(id, balance)
}

func (controller *UserController) DeleteUser(id string) {
    controller.Interactor.DeleteUser(id)
}