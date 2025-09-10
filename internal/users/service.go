package users

type Service interface {
	SignUpUser(username, password, firstName, lastName string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) SignUpUser(username, password, firstName, lastName string) error {
	return s.repository.SaveUser(username, password, firstName, lastName)
}
