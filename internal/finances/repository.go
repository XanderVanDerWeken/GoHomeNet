package finances

import (
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(transactionType TransactionType, money Money, date time.Time, categoryId uint, notes string) error
	GetTransactionsWithYearAndMonth(year, month int) ([]Transaction, error)
}

type CategoryRepository interface {
	SaveCategory(name string) error
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
