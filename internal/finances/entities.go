package finances

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

var (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TransactionType TransactionType `gorm:"not null;index"`
	Money           Money           `gorm:"embedded"`
	Date            time.Time       `gorm:"not null;index"`
	CategoryID      int             `gorm:"not null;index"`
	Notes           string          `gorm:"type:text"`
}

type Category struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name string `gorm:"not null;uniqueIndex"`
}
