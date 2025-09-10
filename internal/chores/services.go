package chores

import (
	"time"

	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Service interface {
	CreateChore(username, title, notes string, dueDate *time.Time) error
	GetChoresByUsername(username string) ([]Chore, error)
	CompleteChore(choreID uint) error
	DeleteChore(choreID uint) error
}

type service struct {
	repository     Repsitory
	userRepository users.Repository
}

func NewService(repository Repsitory, userRepository users.Repository) Service {
	return &service{repository: repository, userRepository: userRepository}
}

func (s *service) CreateChore(username, title, notes string, dueDate *time.Time) error {
	userId, err := s.userRepository.GetUserIdByUsername(username)

	if err != nil {
		return err
	}

	return s.repository.CreateChore(userId, title, notes, dueDate)
}

func (s *service) GetChoresByUsername(username string) ([]Chore, error) {
	return s.repository.GetChoresByUsername(username)
}

func (s *service) CompleteChore(choreID uint) error {
	return s.repository.CompleteChore(choreID)
}

func (s *service) DeleteChore(choreID uint) error {
	return s.repository.DeleteChore(choreID)
}
