package users

import "github.com/xandervanderweken/GoHomeNet/internal/events"

func NewUserRegisteredPersistenceHandler(repository Repository) events.EventHandler {
	return func(e events.Event) {
		if event, ok := e.(UserRegisteredEvent); ok {
			repository.SaveUser(event.Username, event.Password, event.FirstName, event.LastName)
		}
	}
}
