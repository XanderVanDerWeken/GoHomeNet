package recipes

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID      uint   `gorm:"not null;index"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`

	Ingredients  []RecipeIngredient
	Instructions []RecipeStep
}

type RecipeIngredient struct {
	ID       uint `gorm:"primaryKey"`
	RecipeID uint `gorm:"index;not null"`

	Ingredient string `gorm:"not null"`
	Amount     int
	Unit       string `gorm:"not null"`
}

type RecipeStep struct {
	ID       uint `gorm:"primaryKey"`
	RecipeId uint `gorm:"index; not null"`

	Text string `gorm:"not null"`
	Time string `gorm:"not null"`
}
