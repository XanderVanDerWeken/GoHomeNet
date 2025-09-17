package users

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	user1 = User{
		ID:        1,
		Username:  "tester",
		Password:  "123456",
		FirstName: "Test",
		LastName:  "User",
	}
	user2 = User{
		ID:        2,
		Username:  "othertester",
		Password:  "123456",
		FirstName: "Other",
		LastName:  "User",
	}
)

func TestCheckUserCredentials(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	repo := NewRepository(db)

	setupUsers(t, repo)

	invalidUsername := "InvalidUsername"
	invalidPassword := "InvalidPassword"

	// Act
	res1 := repo.CheckUserCredentials(user1.Username, user1.Password)
	res2 := repo.CheckUserCredentials(user2.Username, user2.Password)
	res3 := repo.CheckUserCredentials(invalidUsername, user1.Password)
	res4 := repo.CheckUserCredentials(user1.Username, invalidPassword)
	res5 := repo.CheckUserCredentials(invalidUsername, invalidPassword)

	// Assert
	require.True(t, res1)
	require.True(t, res2)
	require.False(t, res3)
	require.False(t, res4)
	require.False(t, res5)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&User{})

	require.NoError(t, err)
	return db
}

func setupUsers(t *testing.T, repo Repository) {
	var err error
	err = repo.SaveUser(&user1)
	require.NoError(t, err)

	err = repo.SaveUser(&user2)
	require.NoError(t, err)
}
