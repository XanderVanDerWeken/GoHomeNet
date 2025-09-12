package chores

import (
	"time"

	"gorm.io/gorm"
)

type Chore struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID    uint   `gorm:"not null;index"`
	Title     string `gorm:"not null"`
	Notes     string
	DueDate   *time.Time
	Completed bool `gorm:"not null;default:false"`
}
