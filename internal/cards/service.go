package cards

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Service interface {
	AddCard(username string, newCard *Card) error
	GetAllCards() []Card
	GetAllCardsWithUsername(username string) ([]Card, error)
	HandleNewCardEvent(e events.Event)
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

	eventBus.Register("NewCardEvent", s.HandleNewCardEvent)

	return s
}

func (s *service) AddCard(username string, newCard *Card) error {
	userId, err := s.userRepo.GetUserIdByUsername(username)

	if err != nil {
		return shared.ErrUserNotFound
	}

	newCard.UserID = userId

	s.eventBus.Publish(NewCardEvent{
		NewCard: *newCard,
	})
	return nil
}

func (s *service) GetAllCards() []Card {
	return s.repo.GetAllCards()
}

func (s *service) GetAllCardsWithUsername(username string) ([]Card, error) {
	return s.repo.GetAllCardsWithUsername(username)
}

func (s *service) HandleNewCardEvent(e events.Event) {
	if event, ok := e.(NewCardEvent); ok {
		s.repo.AddCard(&event.NewCard)
	}
}
