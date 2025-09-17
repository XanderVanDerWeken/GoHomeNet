package recipes

import (
	"errors"

	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
	"gorm.io/gorm"
)

type Service interface {
	CreateRecipe(username string, newRecipe *Recipe) error
	GetAllRecipes() []Recipe
	GetRecipeWithTitle(title string) (*Recipe, error)
	HandleRecipeCreated(e events.Event)
}

type service struct {
	repo     Repository
	userRepo users.Repository
	eventBus *events.EventBus
}

func NewService(repo Repository, userRepo users.Repository, eventBus *events.EventBus) Service {
	s := &service{
		repo:     repo,
		userRepo: userRepo,
		eventBus: eventBus,
	}

	eventBus.Register("NewRecipeEvent", s.HandleRecipeCreated)

	return s
}

func (s *service) CreateRecipe(username string, newRecipe *Recipe) error {
	var userId uint
	var err error

	if userId, err = s.userRepo.GetUserIdByUsername(username); err != nil {
		return shared.ErrUserNotFound
	}

	if _, err := s.repo.GetRecipeWithTitle(newRecipe.Title); !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecipeAlreadyExists
	}

	newRecipe.UserID = userId

	s.eventBus.Publish(NewRecipeEvent{
		Recipe: *newRecipe,
	})
	return nil
}

func (s *service) GetAllRecipes() []Recipe {
	return s.repo.GetAllRecipes()
}

func (s *service) GetRecipeWithTitle(title string) (*Recipe, error) {
	return s.repo.GetRecipeWithTitle(title)
}

func (s *service) HandleRecipeCreated(e events.Event) {
	if event, ok := e.(NewRecipeEvent); ok {
		s.repo.CreateRecipe(&event.Recipe)
	}
}
