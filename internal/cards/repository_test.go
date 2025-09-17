package cards

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	user1 = users.User{
		ID:        1,
		Username:  "tester",
		Password:  "123456",
		FirstName: "Test",
		LastName:  "User",
	}
	user2 = users.User{
		ID:        2,
		Username:  "othertester",
		Password:  "123456",
		FirstName: "Other",
		LastName:  "User",
	}
	card1 = Card{
		ID:        1,
		UserID:    1,
		Name:      "Card1",
		ExpiresAt: time.Now().Add(31 * 24 * time.Hour),
	}
	card2 = Card{
		ID:        2,
		UserID:    1,
		Name:      "Card2",
		ExpiresAt: time.Now().Add(31 * 24 * time.Hour),
	}
	card3 = Card{
		ID:        3,
		UserID:    2,
		Name:      "Card3",
		ExpiresAt: time.Now().Add(31 * 24 * time.Hour),
	}
)

func TestGetAllCardsWithUsername(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	userRepo := users.NewRepository(db)
	repo := NewRepository(db)

	setupUsers(t, userRepo)
	setupCards(t, repo)

	// Act
	cardsUser1, err1 := repo.GetAllCardsWithUsername(user1.Username)
	cardsUser2, err2 := repo.GetAllCardsWithUsername(user2.Username)
	cardsInvalidUser, err3 := repo.GetAllCardsWithUsername("invalidusername")

	// Assert
	require.NotNil(t, cardsUser1)
	require.Len(t, cardsUser1, 2)
	require.Equal(t, cardsUser1[0].UserID, user1.ID)
	require.Equal(t, cardsUser1[1].UserID, user1.ID)
	require.Equal(t, cardsUser1[0].Name, card1.Name)
	require.Equal(t, cardsUser1[1].Name, card2.Name)
	require.NoError(t, err1)

	require.NotNil(t, cardsUser2)
	require.Len(t, cardsUser2, 1)
	require.Equal(t, cardsUser2[0].UserID, user2.ID)
	require.Equal(t, cardsUser2[0].Name, card3.Name)
	require.NoError(t, err2)

	require.NotNil(t, cardsInvalidUser)
	require.Len(t, cardsInvalidUser, 0)
	require.NoError(t, err3)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&Card{}, &users.User{})

	require.NoError(t, err)
	return db
}

func setupUsers(t *testing.T, repo users.Repository) {
	var err error
	err = repo.SaveUser(&user1)
	require.NoError(t, err)

	err = repo.SaveUser(&user2)
	require.NoError(t, err)
}

func setupCards(t *testing.T, repo Repository) {
	var err error
	err = repo.AddCard(&card1)
	require.NoError(t, err)

	err = repo.AddCard(&card2)
	require.NoError(t, err)

	err = repo.AddCard(&card3)
	require.NoError(t, err)
}
