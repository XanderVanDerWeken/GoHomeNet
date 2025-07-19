package chores

import "time"

type Chore struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description *string   `gorm:"null"`
	DueDate     time.Time `gorm:"not null"`
	IsDone      bool      `gorm:"default:false"`
}
