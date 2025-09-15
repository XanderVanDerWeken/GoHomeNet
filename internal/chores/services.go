package chores

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Service interface {
	CreateChore(username string, newChore *Chore) error
	GetAllChores() []Chore
	GetChoresByUsername(username string) ([]Chore, error)
	CompleteChore(choreID uint) error
	DeleteChore(choreID uint) error

	HandleNewChoreEvent(e events.Event)
	HandleCompletedChoreEvent(e events.Event)
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

func (s *service) CreateChore(username string, newChore *Chore) error {
	userId, err := s.userRepo.GetUserIdByUsername(username)

	if err != nil {
		return shared.ErrUserNotFound
	}

	newChore.UserID = userId
	s.eventBus.Publish(NewChoreEvent{
		NewChore: *newChore,
	})

	return nil
}

func (s *service) GetAllChores() []Chore {
	return s.repo.GetAllChores()
}

func (s *service) GetChoresByUsername(username string) ([]Chore, error) {
	return s.repo.GetChoresByUsername(username)
}

func (s *service) CompleteChore(choreID uint) error {
	return s.repo.CompleteChore(choreID)
}

func (s *service) DeleteChore(choreID uint) error {
	return s.repo.DeleteChore(choreID)
}

func (s *service) HandleNewChoreEvent(e events.Event) {
	if event, ok := e.(NewChoreEvent); ok {
		s.repo.CreateChore(&event.NewChore)
	}
}

func (s *service) HandleCompletedChoreEvent(e events.Event) {
	if event, ok := e.(CompletedChoreEvent); ok {
		s.repo.CompleteChore(event.ChoreId)
	}
}
