package recipes

import "gorm.io/gorm"

type Repository interface {
	CreateRecipe(newRecipe *Recipe) error
	GetAllRecipes() []Recipe
	GetRecipeWithTitle(title string) (*Recipe, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateRecipe(newRecipe *Recipe) error {

	if err := r.db.Create(newRecipe).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllRecipes() []Recipe {
	var recipes []Recipe

	r.db.
		Preload("Ingredients").
		Preload("Instructions").
		Find(&recipes)

	return recipes
}

func (r *repository) GetRecipeWithTitle(title string) (*Recipe, error) {
	var recipe Recipe
	err := r.db.
		Where("title = ?", title).
		First(&recipe).Error

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
