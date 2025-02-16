package model

type Item struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"unique;not null"`
	Price int    `gorm:"default:1000"`
}
