package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name"`
    Balance int64 `json:"balance"`
    //Moneys []*Money `json:"money"`
}
