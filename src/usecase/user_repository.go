package usecase

import (
    "money-send-api/src/domain"
)

// usecase/user_interactorが引数として受けるために必要
type UserRepository interface {
    AddUser(domain.User)
    GetAllUsers() []domain.User
    GetUser(string) domain.User
    UpdateAllBalance(balance int64)
    UpdateBalance(id string, balance int64)
    DeleteUser(string)
}
