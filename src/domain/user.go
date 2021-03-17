package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name"`
    Balance int64 `json:"balance"`
    IsBalanceReceivable bool `json:"is_balance_receivable"`
    //Moneys []*Money `json:"money"`
}
