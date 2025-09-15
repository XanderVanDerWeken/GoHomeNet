package finances

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateAggregation(t *testing.T) {
	// Arrange
	mockRepo := mockCategoryRepository{
		categories: map[uint]Category{
			1: {ID: 1, Name: "Work"},
			2: {ID: 2, Name: "Food"},
		},
	}
	service := &service{categoryRepo: &mockRepo}

	transactions := []Transaction{
		{TransactionType: TransactionTypeIncome, Money: NewMoney(100, 0), CategoryID: 1},
		{TransactionType: TransactionTypeExpense, Money: NewMoney(50, 0), CategoryID: 1},
		{TransactionType: TransactionTypeExpense, Money: NewMoney(20, 0), CategoryID: 2},
	}

	year := 2025
	month := 6

	// Act
	result, err := service.calculateAggregation(year, month, transactions)

	// Assert
	require.NoError(t, err)

	require.Equal(t, result.Year, year)
	require.Equal(t, result.Month, month)

	require.Len(t, result.Transactions, 2)

	require.Equal(t, result.TotalIncome.Cents, int64(10000))
	require.Equal(t, result.TotalExpense.Cents, int64(7000))

	names := []string{result.Transactions[0].CategoryName, result.Transactions[1].CategoryName}
	require.Contains(t, names, "Work")
	require.Contains(t, names, "Food")
}

type mockCategoryRepository struct {
	categories map[uint]Category
}

func (m *mockCategoryRepository) GetCategoryById(id uint) (*Category, error) {
	category := m.categories[id]
	return &category, nil
}

func (m *mockCategoryRepository) GetCategoryByName(name string) (*Category, error) {
	return &Category{}, nil
}

func (m *mockCategoryRepository) SaveCategory(newCategory *Category) error {
	return nil
}

func (m *mockCategoryRepository) GetAllCategories() []Category {
	return nil
}
