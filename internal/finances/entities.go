package finances

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name string `gorm:"not null;uniqueIndex"`
}
