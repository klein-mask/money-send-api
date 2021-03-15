package domain

type Money struct {
    ID   int    `json:"id" gorm:"primary_key"`
	User User
	CreatedAt time.Time
	UpdatedAt time.Time
    HistoryId string `json:"history_id" gorm:"unique;not null;index:idx_history_id"`
}
