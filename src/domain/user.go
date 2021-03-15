package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name"`
    Balance uint `json:"balance"`
    //Moneys []*Money `json:"money"`
}
