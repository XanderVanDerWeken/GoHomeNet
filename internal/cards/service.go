package cards

import (
	"time"

	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Service interface {
	AddCard(username, name string, expiresAt time.Time) error
	GetAllCards() []Card
	GetAllOwnCards(username string) ([]Card, error)
}

type service struct {
	repository     Repository
	userRepository users.Repository
}

func NewService(repository Repository, userRepository users.Repository) Service {
	return &service{repository: repository, userRepository: userRepository}
}

func (s *service) AddCard(username, name string, expiresAt time.Time) error {
	userId, err := s.userRepository.GetUserIdByUsername(username)

	if err != nil {
		return err
	}

	return s.repository.AddCard(userId, name, expiresAt)
}

func (s *service) GetAllCards() []Card {
	return s.repository.GetAllCards()
}

func (s *service) GetAllOwnCards(username string) ([]Card, error) {
	return s.repository.GetAllOwnCards(username)
}
