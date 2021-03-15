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
    db.Where("id = ?", id).Find(&user)
    return user
}

func (db *UserRepository) Delete(id string) {
    user := []domain.User{}
    db.DeleteById(&user, id)
}
