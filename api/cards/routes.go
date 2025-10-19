package cards

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/cards"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func Routes(router chi.Router, service cards.Service, userService users.Service) {
	handler := NewCardHandler(service, userService)

	router.Post("/", handler.PostNewCard)
	router.Get("/", handler.GetCards)
}
