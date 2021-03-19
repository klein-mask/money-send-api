package database

import (
    "net/http"
    "money-send-api/domain"
    _ "errors"
    _ "fmt"
)

type UserRepository struct {
    SqlHandler
}

type UserError struct {
    Code int 
    Message string
}

func (e UserError) Error() string {
    em := "[Code]: " + string(e.Code) + ", [Message]: " + e.Message
    return em
}

func (db *UserRepository) Login(name string, password string) error {
    user := domain.User{}
    err := db.FindByName(&user, name)
    if err != nil {
        return &UserError{http.StatusInternalServerError, err.Error()}
    }

    if user.Name == "" || user.Name != name || user.Password == "" || user.Password != password {
        return &UserError{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
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