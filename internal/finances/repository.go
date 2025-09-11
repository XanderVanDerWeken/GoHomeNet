package finances

import (
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(transactionType TransactionType, money Money, date time.Time, categoryId uint, notes string) error
}

type CategoryRepository interface {
	SaveCategory(name string) error
	GetAllCategories() []Category
	GetCategoryByName(name string) (*Category, error)
}

type transactionRepository struct {
	db *gorm.DB
}

type categoryRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *transactionRepository) SaveTransaction(transactionType TransactionType, money Money, date time.Time, categoryId uint, notes string) error {
	transaction := Transaction{
		TransactionType: transactionType,
		Money:           money,
		Date:            date,
		CategoryID:      categoryId,
		Notes:           notes,
	}

	return r.db.Create(&transaction).Error
}

func (r *categoryRepository) SaveCategory(name string) error {
	category := Category{
		Name: name,
	}

	return r.db.Create(&category).Error
}

func (r *categoryRepository) GetAllCategories() []Category {
	var categories []Category
	r.db.Find(&categories)

	return categories
}

func (r *categoryRepository) GetCategoryByName(name string) (*Category, error) {
	var category Category
	err := r.db.Where("name = ?", name).First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
