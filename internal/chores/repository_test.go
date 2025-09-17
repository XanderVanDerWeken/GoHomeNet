package chores

import (
	"testing"

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
	chore1 = Chore{
		ID:        1,
		UserID:    1,
		Title:     "Chore1",
		Notes:     "Chore1 Notes",
		Completed: false,
	}
	chore2 = Chore{
		ID:        2,
		UserID:    1,
		Title:     "Chore2",
		Notes:     "Chore2 Notes",
		Completed: false,
	}
	chore3 = Chore{
		ID:        3,
		UserID:    2,
		Title:     "Chore3",
		Notes:     "Chore3 Notes",
		Completed: false,
	}
)

func TestGetChoresByUsername(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	userRepo := users.NewRepository(db)
	repo := NewRepository(db)

	setupUsers(t, userRepo)
	setupChores(t, repo)

	// Act
	choresUser1, err1 := repo.GetChoresByUsername(user1.Username)
	choresUser2, err2 := repo.GetChoresByUsername(user2.Username)
	choresInvalidUser, err3 := repo.GetChoresByUsername("InvalidUser")

	// Assert
	require.NotNil(t, choresUser1)
	require.Len(t, choresUser1, 2)
	require.Equal(t, choresUser1[0].UserID, user1.ID)
	require.Equal(t, choresUser1[1].UserID, user1.ID)
	require.Equal(t, choresUser1[0].Title, chore1.Title)
	require.Equal(t, choresUser1[1].Title, chore2.Title)
	require.NoError(t, err1)

	require.NotNil(t, choresUser2)
	require.Len(t, choresUser2, 1)
	require.Equal(t, choresUser2[0].UserID, user2.ID)
	require.Equal(t, choresUser2[0].Title, chore3.Title)
	require.NoError(t, err2)

	require.NotNil(t, choresInvalidUser)
	require.Len(t, choresInvalidUser, 0)
	require.NoError(t, err3)
}

func TestGetChoreById(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	userRepo := users.NewRepository(db)
	repo := NewRepository(db)

	setupUsers(t, userRepo)
	setupChores(t, repo)

	// Act
	foundChore1, err1 := repo.GetChoreById(chore1.ID)
	foundChore2, err2 := repo.GetChoreById(chore2.ID)
	notFoundChore, err3 := repo.GetChoreById(999)

	// Assert
	require.NotNil(t, foundChore1)
	require.NoError(t, err1)

	require.NotNil(t, foundChore2)
	require.NoError(t, err2)

	require.Nil(t, notFoundChore)
	require.Error(t, err3)
	require.ErrorIs(t, err3, gorm.ErrRecordNotFound)
}

func TestCompleteChore(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	userRepo := users.NewRepository(db)
	repo := NewRepository(db)

	setupUsers(t, userRepo)
	setupChores(t, repo)

	// Act
	isDoneBefore := chore1.Completed
	err := repo.CompleteChore(chore1.ID)

	// Assert
	require.False(t, isDoneBefore)
	require.NoError(t, err)

	updatedChore, err := repo.GetChoreById(chore1.ID)
	require.NoError(t, err)
	require.NotNil(t, updatedChore)
	require.True(t, updatedChore.Completed)
}

func TestDeleteChore(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	userRepo := users.NewRepository(db)
	repo := NewRepository(db)

	setupUsers(t, userRepo)
	setupChores(t, repo)

	// Act
	err1 := repo.DeleteChore(chore2.ID)
	err2 := repo.DeleteChore(999)

	// Assert
	require.NoError(t, err1)
	deletedChore, err3 := repo.GetChoreById(chore2.ID)
	require.Error(t, err3)
	require.ErrorIs(t, err3, gorm.ErrRecordNotFound)
	require.Nil(t, deletedChore)

	require.NoError(t, err2)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&Chore{}, &users.User{})

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

func setupChores(t *testing.T, repo Repository) {
	var err error
	err = repo.CreateChore(&chore1)
	require.NoError(t, err)

	err = repo.CreateChore(&chore2)
	require.NoError(t, err)

	err = repo.CreateChore(&chore3)
	require.NoError(t, err)
}
