package model

type Buy struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	ItemID uint `gorm:"not null"`
	Count  uint `gorm:"default:1"`
}
