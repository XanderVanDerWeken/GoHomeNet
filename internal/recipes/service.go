package recipes

import (
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Service interface {
	CreateRecipe(username, title string) error
	GetAllRecipes() []Recipe
	GetRecipeWithTitle(title string) (*Recipe, error)
}

type service struct {
	repo     Repository
	userRepo users.Repository
}

func NewService(repo Repository, userRepo users.Repository) Service {
	return &service{repo: repo, userRepo: userRepo}
}

func (s *service) CreateRecipe(username, title string) error {
	userId, err := s.userRepo.GetUserIdByUsername(username)

	if err != nil {
		return shared.ErrUserNotFound
	}

	return s.repo.CreateRecipe(userId, title)
}

func (s *service) GetAllRecipes() []Recipe {
	return s.repo.GetAllRecipes()
}

func (s *service) GetRecipeWithTitle(title string) (*Recipe, error) {
	return s.repo.GetRecipeWithTitle(title)
}
