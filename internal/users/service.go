package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Authenticate(email, password string) (*User, error)
	GetUserByID(id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Authenticate(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}
