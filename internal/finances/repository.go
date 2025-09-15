package finances

import (
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(newTransaction *Transaction) error
	GetTransactionsWithYearAndMonth(year, month int) ([]Transaction, error)
}

type CategoryRepository interface {
	SaveCategory(newCategory *Category) error
	GetAllCategories() []Category
	GetCategoryByName(name string) (*Category, error)
	GetCategoryById(id uint) (*Category, error)
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

func (r *transactionRepository) SaveTransaction(newTransaction *Transaction) error {
	return r.db.Create(newTransaction).Error
}

func (r *categoryRepository) SaveCategory(newCategory *Category) error {
	return r.db.Create(newCategory).Error
}

func (r *transactionRepository) GetTransactionsWithYearAndMonth(year, month int) ([]Transaction, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	var transactions []Transaction

	err := r.db.Where("date >= ? AND date < ?", startDate, endDate).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
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

func (r *categoryRepository) GetCategoryById(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
