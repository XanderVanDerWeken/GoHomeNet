package users

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type Service interface {
	SignUpUser(username, password, firstName, lastName string)
}

type service struct {
	eventBus *events.EventBus
}

func NewService(eventBus *events.EventBus) Service {
	return &service{eventBus: eventBus}
}

func (s *service) SignUpUser(username, password, firstName, lastName string) {
	s.eventBus.Publish(UserRegisteredEvent{
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	})
}
