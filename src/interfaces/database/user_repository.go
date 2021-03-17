package database

import (
    "money-send-api/src/domain"
)

type UserRepository struct {
    SqlHandler
}

func (db *UserRepository) AddUser(u domain.User) {
    db.Create(&u)
}

func (db *UserRepository) GetAllUsers() []domain.User {
    users := []domain.User{}
    db.FindAll(&users)
    return users
}

func (db *UserRepository) GetUser(id string) domain.User {
    user := domain.User{}
    db.FindById(&user, id)
    return user
}

func (db *UserRepository) UpdateAllBalance(balance int64) {
    user := domain.User{}
    balanceExpr := []int64{balance}
    db.UpdateByExpr(&user, "is_balance_receivable = ?", true, "balance", "balance + ?", balanceExpr)
}

func (db *UserRepository) UpdateBalance(id string, balance int64) {
    user := domain.User{}
    db.Update(&user, "ID = ?", id, "balance", balance)
}

func (db *UserRepository) DeleteUser(id string) {
    user := []domain.User{}
    db.DeleteById(&user, id)
}
