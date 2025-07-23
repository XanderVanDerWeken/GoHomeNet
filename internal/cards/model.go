package cards

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Name    string    `gorm:"not null"`
	DueDate time.Time `gorm:"not null"`
}
