package recipes

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeltedAt  gorm.DeletedAt `gorm:"index"`

	UserID uint   `gorm:"not null;index"`
	Title  string `gorm:"not null"`
}
