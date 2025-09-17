package auth

import "github.com/xandervanderweken/GoHomeNet/internal/users"

type Service interface {
	SignUpUser(username, password, firstName, lastName string) error
}

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) Service {
	return &service{repository: repository}
}

func (s *service) SignUpUser(username, password, firstName, lastName string) error {
	return s.repository.SaveUser(username, password, firstName, lastName)
}
