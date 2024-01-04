package models

type User struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
