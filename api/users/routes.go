package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func Routes(router chi.Router, service users.Service) {
	handler := NewUserHandler(service)

	router.Get("/{username}", handler.GetUserByUsername)
}
