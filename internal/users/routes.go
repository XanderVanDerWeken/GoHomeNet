package users

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router, service Service) {
	handler := UserHandler{service: service}

	router.Post("/signup", handler.PostSignupUser)
}
