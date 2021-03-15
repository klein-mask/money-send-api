package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    //ID   int    `json:"id" gorm:"primary_key"`
    Name string `json:"name"`
    MoneyId  string `json:"money_id" gorm:"unique;not null;index:idx_money_id"`
    //Password string `json:"password"`
}
