package finances

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Category1 = Category{
		ID:   1,
		Name: "Category1",
	}
	Category2 = Category{
		ID:   2,
		Name: "Category2",
	}
	Transaction1 = Transaction{
		ID:              1,
		TransactionType: TransactionTypeIncome,
		Money:           NewMoneyFromCents(1000),
		Date:            time.Now(),
		CategoryID:      Category1.ID,
		Notes:           "Transaction1 Notes",
	}
	Transaction2 = Transaction{
		ID:              2,
		TransactionType: TransactionTypeExpense,
		Money:           NewMoneyFromCents(500),
		Date:            time.Now(),
		CategoryID:      Category2.ID,
		Notes:           "Transaction2 Notes",
	}
	Transaction3 = Transaction{
		ID:              3,
		TransactionType: TransactionTypeExpense,
		Money:           NewMoneyFromCents(550),
		Date:            time.Now(),
		CategoryID:      Category2.ID,
		Notes:           "Transaction3 Notes",
	}
)

func TestGetCategoryByName(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	categoryRepo := NewCategoryRepository(db)
	transactionRepo := NewTransactionRepository(db)

	setupCategories(t, categoryRepo)
	setupTransactions(t, transactionRepo)

	// Act
	category1, err1 := categoryRepo.GetCategoryByName(Category1.Name)
	category2, err2 := categoryRepo.GetCategoryByName(Category2.Name)
	categoryInvalid, err3 := categoryRepo.GetCategoryByName("InvalidCategory")

	// Assert
	require.NotNil(t, category1)
	require.Equal(t, category1.Name, Category1.Name)
	require.Equal(t, category1.ID, Category1.ID)
	require.NoError(t, err1)

	require.NotNil(t, category2)
	require.Equal(t, category2.Name, Category2.Name)
	require.Equal(t, category2.ID, Category2.ID)
	require.NoError(t, err2)

	require.Nil(t, categoryInvalid)
	require.Error(t, err3)
	require.ErrorIs(t, err3, gorm.ErrRecordNotFound)
}

func TestGetCategoryById(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	categoryRepo := NewCategoryRepository(db)
	transactionRepo := NewTransactionRepository(db)

	setupCategories(t, categoryRepo)
	setupTransactions(t, transactionRepo)

	// Act
	category1, err1 := categoryRepo.GetCategoryById(Category1.ID)
	category2, err2 := categoryRepo.GetCategoryById(Category2.ID)
	categoryInvalid, err3 := categoryRepo.GetCategoryById(999)

	// Assert
	require.NotNil(t, category1)
	require.Equal(t, category1.Name, Category1.Name)
	require.Equal(t, category1.ID, Category1.ID)
	require.NoError(t, err1)

	require.NotNil(t, category2)
	require.Equal(t, category2.Name, Category2.Name)
	require.Equal(t, category2.ID, Category2.ID)
	require.NoError(t, err2)

	require.Nil(t, categoryInvalid)
	require.Error(t, err3)
	require.ErrorIs(t, err3, gorm.ErrRecordNotFound)
}

func TestGetTransactionsWithYearAndMonth(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	categoryRepo := NewCategoryRepository(db)
	transactionRepo := NewTransactionRepository(db)

	setupCategories(t, categoryRepo)
	setupTransactions(t, transactionRepo)

	// Act

	// Assert
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&Transaction{}, &Category{})

	require.NoError(t, err)
	return db
}

func setupCategories(t *testing.T, repo CategoryRepository) {
	err1 := repo.SaveCategory(&Category1)
	err2 := repo.SaveCategory(&Category2)

	require.NoError(t, err1)
	require.NoError(t, err2)
}

func setupTransactions(t *testing.T, repo TransactionRepository) {
	err1 := repo.SaveTransaction(&Transaction1)
	err2 := repo.SaveTransaction(&Transaction2)
	err3 := repo.SaveTransaction(&Transaction3)

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)
}
