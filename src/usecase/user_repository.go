package usecase

import (
    "money-send-api/src/domain"
)

// usecase/user_interactorが引数として受けるために必要
type UserRepository interface {
    Store(domain.User)
    Select() []domain.User
    SelectUser(string) domain.User
    UpdateBalance(id string, balance int64)
    Delete(string)
}
