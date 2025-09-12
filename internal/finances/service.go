package finances

import "time"

type Service interface {
	CreateCategory(name string) error
	GetAllCategories() []Category

	CreateTransaction(transactionType TransactionType, money Money, date time.Time, categoryName, notes string) error
	GetAggregatedTransactions(year, month int) (*AggregationResult, error)
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

func (s *service) GetAggregatedTransactions(year, month int) (*AggregationResult, error) {
	transactions, err := s.transactionRepository.GetTransactionsWithYearAndMonth(year, month)

	if err != nil {
		return nil, err
	}

	return s.calculateAggregation(year, month, transactions)
}

func (s *service) calculateAggregation(year, month int, transactions []Transaction) (*AggregationResult, error) {
	result := AggregationResult{
		Year:  year,
		Month: month,
	}

	aggregationMap := make(map[uint]*AggregatedTransaction)

	for _, tx := range transactions {
		switch tx.TransactionType {
		case TransactionTypeIncome:
			result.TotalIncome.Add(tx.Money)
		case TransactionTypeExpense:
			result.TotalExpense.Add(tx.Money)
		}

		if aggregationMap[tx.CategoryID] == nil {
			category, err := s.categoryRepository.GetCategoryById(tx.CategoryID)

			if err != nil {
				return nil, err
			}

			aggregationMap[tx.CategoryID] = &AggregatedTransaction{
				CategoryName:    category.Name,
				Money:           Money{},
				TransactionType: tx.TransactionType,
			}
			aggregationMap[tx.CategoryID].Money.Add(tx.Money)
		}
	}

	for _, agg := range aggregationMap {
		result.Transactions = append(result.Transactions, *agg)
	}

	return &result, nil
}
