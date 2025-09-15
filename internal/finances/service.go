package finances

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
)

type Service interface {
	CreateCategory(newCategory *Category) error
	GetAllCategories() []Category

	CreateTransaction(categoryName string, newTransaction *Transaction) error
	GetAggregatedTransactions(year, month int) (*AggregationResult, error)

	HandleNewCategoryEvent(e events.Event)
	HandleNewTransactionEvent(e events.Event)
}

type service struct {
	transactionRepo TransactionRepository
	categoryRepo    CategoryRepository
	eventBus        *events.EventBus
}

func NewService(transactionRepo TransactionRepository, categoryRepo CategoryRepository, eventBus *events.EventBus) Service {
	return &service{
		transactionRepo: transactionRepo,
		categoryRepo:    categoryRepo,
		eventBus:        eventBus,
	}
}

func (s *service) CreateCategory(newCategory *Category) error {
	if cat, err := s.categoryRepo.GetCategoryByName(newCategory.Name); err != nil {
		return err
	} else if cat != nil {
		return ErrCategoryAlreadyExists
	}

	s.eventBus.Publish(NewCategoryEvent{
		NewCategory: *newCategory,
	})

	return nil
}

func (s *service) GetAllCategories() []Category {
	return s.categoryRepo.GetAllCategories()
}

func (s *service) CreateTransaction(categoryName string, newTransaction *Transaction) error {
	category, err := s.categoryRepo.GetCategoryByName(categoryName)

	if err != nil {
		return err
	}

	newTransaction.CategoryID = category.ID

	s.eventBus.Publish(NewTransactionEvent{
		NewTransaction: *newTransaction,
	})

	return nil
}

func (s *service) GetAggregatedTransactions(year, month int) (*AggregationResult, error) {
	transactions, err := s.transactionRepo.GetTransactionsWithYearAndMonth(year, month)

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
			category, err := s.categoryRepo.GetCategoryById(tx.CategoryID)

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

func (s *service) HandleNewCategoryEvent(e events.Event) {
	if event, ok := e.(NewCategoryEvent); ok {
		s.categoryRepo.SaveCategory(&event.NewCategory)
	}
}

func (s *service) HandleNewTransactionEvent(e events.Event) {
	if event, ok := e.(NewTransactionEvent); ok {
		s.transactionRepo.SaveTransaction(&event.NewTransaction)
	}
}
