package usecase

import (
    "money-send-api/domain"
)

// usecase/user_interactorが引数として受けるために必要
type UserRepository interface {
    AddUser(domain.User) error
    GetAllUsers() ([]domain.User, error)
    GetUser(string) (domain.User, error)
    UpdateAllBalance(balance int64) error
    UpdateBalance(id string, balance int64) error
    DeleteAllUser() error
    DeleteUser(string) error
}
