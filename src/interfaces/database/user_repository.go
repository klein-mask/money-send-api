package database

import (
    "money-send-api/domain"
    "errors"
)

type UserRepository struct {
    SqlHandler
}

func (db *UserRepository) Login(name string, password string) error {
    user := domain.User{}
    err := db.FindByName(&user, name)
    if err != nil {
        return err
    }
    if user.Name != name || user.Password != password {
        return errors.New("Login faild.")
    }

    return nil
}

func (db *UserRepository) Regist(u domain.User) error {
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
    balanceExpr := []int64{balance}
    return db.UpdateByExpr(&user, "ID = ?", userId, "balance", "balance + ?", balanceExpr)
}

func (db *UserRepository) DeleteUser(id string) error {
    user := []domain.User{}
    return db.DeleteById(&user, id)
}