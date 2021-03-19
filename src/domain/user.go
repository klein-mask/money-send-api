package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name"`
    Password string `json:"password"`
    Balance int64 `json:"balance"`
    IsBalanceReceivable bool `json:"is_balance_receivable"`
}
