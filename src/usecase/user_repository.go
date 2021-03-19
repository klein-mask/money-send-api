package usecase

import (
    "money-send-api/domain"
)

type UserRepository interface {
    Login(name string, password string) error
    Regist(domain.User) error
    GetAllUsers() ([]domain.User, error)
    GetUser(id string) (domain.User, error)
    UpdateAllBalance(balance int64) error
    UpdateBalance(id string, balance int64) error
    DeleteUser(id string) error
}
