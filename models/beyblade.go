package models

type Beyblade struct {
	ID               uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string  `json:"name"`
	Damage           uint    `json:"damage"`
	Shield           uint    `json:"shield"`
	SpecialMonsterID uint    `json:"SpecialMonsterID"`
	Monster          Monster `gorm:"foreignKey:SpecialMonsterID" json:"-"`
	OwnerID          uint    `json:"owner_id"`
	User             User    `gorm:"foreignKey:OwnerID" json:"-"`
}
