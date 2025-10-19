package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/auth"
)

func Routes(router chi.Router, service auth.Service) {
	handler := NewAuthHandler(service)

	router.Post("/signup", handler.PostSignupUser)
	router.Post("/login", handler.PostLoginUser)
}
