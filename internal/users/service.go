package users

type Service interface {
	GetUserByUserId(userId uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetUserByUserId(userId uint) (*User, error) {
	return s.repository.GetUserByUserId(userId)
}

func (s *service) GetUserByUsername(username string) (*User, error) {
	return s.repository.GetUserByUsername(username)
}
