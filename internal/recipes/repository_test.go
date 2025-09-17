package recipes

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	recipe1 = Recipe{
		Title:       "Recipe1",
		Description: "Recipe1",
		Ingredients: []RecipeIngredient{
			{
				Ingredient: "Ingredient1",
				Amount:     1,
				Unit:       "Unit",
			},
		},
		Instructions: []RecipeStep{
			{
				Text: "Step1",
				Time: "TimeAmount",
			},
			{
				Text: "Step2",
				Time: "TimeAmount",
			},
		},
	}
	recipe2 = Recipe{
		Title:       "Recipe2",
		Description: "Recipe2",
		Ingredients: []RecipeIngredient{
			{
				Ingredient: "Ingredient2",
				Amount:     2,
				Unit:       "Unit",
			},
		},
		Instructions: []RecipeStep{
			{
				Text: "Step3",
				Time: "TimeAmount",
			},
		},
	}
)

func TestGetAllRecipes(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	repo := NewRepository(db)

	setupRecipes(t, repo)

	// Act
	recipes := repo.GetAllRecipes()

	// Assert
	require.Len(t, recipes, 2)

	require.Equal(t, recipes[0].Title, recipe1.Title)
	require.Equal(t, recipes[0].Description, recipe1.Description)
	require.Len(t, recipes[0].Ingredients, len(recipe1.Ingredients))
	require.Len(t, recipes[0].Instructions, len(recipe1.Instructions))

	require.Equal(t, recipes[1].Title, recipe2.Title)
	require.Equal(t, recipes[1].Description, recipe2.Description)
	require.Len(t, recipes[1].Ingredients, len(recipe2.Ingredients))
	require.Len(t, recipes[1].Instructions, len(recipe2.Instructions))
}

func TestGetRecipeWithTitle(t *testing.T) {
	// Arrange
	db := setupTestDB(t)
	repo := NewRepository(db)

	setupRecipes(t, repo)

	// Act
	foundRecipe1, err1 := repo.GetRecipeWithTitle(recipe1.Title)
	foundRecipe2, err2 := repo.GetRecipeWithTitle(recipe2.Title)
	notFoundRecipe, err3 := repo.GetRecipeWithTitle("InvalidRecipe")

	// Assert
	require.NotNil(t, foundRecipe1)
	require.Equal(t, foundRecipe1.Title, recipe1.Title)
	require.Equal(t, foundRecipe1.Description, recipe1.Description)
	require.Len(t, foundRecipe1.Ingredients, len(recipe1.Ingredients))
	require.Len(t, foundRecipe1.Instructions, len(recipe1.Instructions))
	require.NoError(t, err1)

	require.NotNil(t, foundRecipe2)
	require.Equal(t, foundRecipe2.Title, recipe2.Title)
	require.Equal(t, foundRecipe2.Description, recipe2.Description)
	require.Len(t, foundRecipe2.Ingredients, len(recipe2.Ingredients))
	require.Len(t, foundRecipe2.Instructions, len(recipe2.Instructions))
	require.NoError(t, err2)

	require.Nil(t, notFoundRecipe)
	require.Error(t, err3)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&users.User{}, &Recipe{}, &RecipeStep{}, &RecipeIngredient{})
	require.NoError(t, err)
	return db
}

func setupRecipes(t *testing.T, repo Repository) {
	var err error
	err = repo.CreateRecipe(&recipe1)
	require.NoError(t, err)

	err = repo.CreateRecipe(&recipe2)
	require.NoError(t, err)
}
