package auth

import "github.com/go-chi/chi/v5"

func Routes(router chi.Router, service Service) {
	handler := NewAuthHandler(service)

	router.Post("/signup", handler.PostSignupUser)
	router.Post("/login", handler.PostLoginUser)
}
