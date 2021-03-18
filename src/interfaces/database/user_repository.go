package database

import (
    "money-send-api/domain"
)

type UserRepository struct {
    SqlHandler
}

func (db *UserRepository) AddUser(u domain.User) error {
    return db.Create(&u)
}

func (db *UserRepository) GetAllUsers() ([]domain.User, error) {
    users := []domain.User{}
    err := db.FindAll(&users)
    return users, err
}

func (db *UserRepository) GetUser(id string) (domain.User, error) {
    user := domain.User{}
    err := db.FindById(&user, id)
    return user, err
}

func (db *UserRepository) UpdateAllBalance(balance int64) error {
    user := domain.User{}
    balanceExpr := []int64{balance}
    return db.UpdateByExpr(&user, "is_balance_receivable = ?", true, "balance", "balance + ?", balanceExpr)
}

func (db *UserRepository) UpdateBalance(userId string, balance int64) error {
    user := domain.User{}
    return db.Update(&user, "ID = ?", userId, "balance", balance)
}

func (db *UserRepository) DeleteUser(id string) error {
    user := []domain.User{}
    return db.DeleteById(&user, id)
}