package chores

import (
	"time"

	"gorm.io/gorm"
)

type Chore struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description *string   `gorm:"null"`
	DueDate     time.Time `gorm:"not null"`
	IsDone      bool      `gorm:"default:false"`
}
