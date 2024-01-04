package models

type Monster struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `json:"Name"`
	Specialattack string `json:"specialattack"`
	Elements      string `json:"elements"`
}
