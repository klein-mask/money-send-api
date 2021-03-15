package domain

import (
    "gorm.io/gorm"
)

type Money struct {
    gorm.Model
    Balance uint `json:"balance"`
	UserId uint `json:"user_id" gorm:"unique;not null"`
    HistoryId string `json:"history_id" gorm:"unique;not null"`
}
