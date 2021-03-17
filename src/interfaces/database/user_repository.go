package database

import (
    "money-send-api/src/domain"
)

type UserRepository struct {
    SqlHandler
}

func (db *UserRepository) Store(u domain.User) {
    db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
    users := []domain.User{}
    db.FindAll(&users)
    return users
}

func (db *UserRepository) SelectUser(id string) domain.User {
    user := domain.User{}
    db.FindById(&user, id)
    return user
}

func (db *UserRepository) UpdateBalance(id string, balance int64) {
    user := domain.User{}
    db.UpdateBalanceById(&user, id, balance)
}

func (db *UserRepository) Delete(id string) {
    user := []domain.User{}
    db.DeleteById(&user, id)
}
