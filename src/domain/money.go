package domain

import (
    "gorm.io/gorm"
)

type Money struct {
    gorm.Model
	User User
    HistoryId string `json:"history_id" gorm:"unique;not null"`
}
