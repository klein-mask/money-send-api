package usecase

import (
    "money-send-api/src/domain"
)

type UserRepository interface {
    Store(domain.User)
    Select() []domain.User
    Delete(id string)
}
