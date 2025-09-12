package chores

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func Routes(router chi.Router, service Service, userService users.Service) {
	handler := NewChoreHandler(service, userService)

	router.Post("/", handler.PostNewChore)
	router.Get("/", handler.GetChores)

	router.Put("/{choreID}/complete", handler.PutChoreComplete)
	router.Delete("/{choreID}", handler.DeleteChore)
}
