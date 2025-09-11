package finances

import "time"

type Service interface {
	CreateCategory(name string) error
	GetAllCategories() []Category

	CreateTransaction(transactionType TransactionType, money Money, date time.Time, categoryName, notes string) error
}

type service struct {
	transactionRepository TransactionRepository
	categoryRepository    CategoryRepository
}

func NewService(transactionRepository TransactionRepository, catcategoryRepository CategoryRepository) Service {
	return &service{
		transactionRepository: transactionRepository,
		categoryRepository:    catcategoryRepository,
	}
}

func (s *service) CreateCategory(name string) error {
	return s.categoryRepository.SaveCategory(name)
}

func (s *service) GetAllCategories() []Category {
	return s.categoryRepository.GetAllCategories()
}

func (s *service) CreateTransaction(transactionType TransactionType, money Money, date time.Time, categoryName, notes string) error {
	category, err := s.categoryRepository.GetCategoryByName(categoryName)

	if err != nil {
		return err
	}

	return s.transactionRepository.SaveTransaction(transactionType, money, date, category.ID, notes)
}
