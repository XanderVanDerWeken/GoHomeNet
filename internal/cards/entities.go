package cards

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID    uint      `gorm:"not null;index"`
	Name      string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
