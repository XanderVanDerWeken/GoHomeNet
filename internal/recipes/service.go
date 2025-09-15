package recipes

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
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
	return &service{
		repo:     repo,
		userRepo: userRepo,
		eventBus: eventBus,
	}
}

func (s *service) CreateRecipe(username string, newRecipe *Recipe) error {
	userId, err := s.userRepo.GetUserIdByUsername(username)

	if err != nil {
		return shared.ErrUserNotFound
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
