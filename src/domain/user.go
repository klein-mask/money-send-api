package domain

type User struct {
    ID   int    `json:"id" gorm:"primary_key"`
    Name string `json:"name"`
    Password string `json:"password"`
}
