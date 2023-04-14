package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
