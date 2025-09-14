package recipes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func Routes(router chi.Router, service Service, userService users.Service) {
	handler := NewRecipeHandler(service, userService)

	router.Post("/", handler.PostNewRecipe)
	router.Get("/", handler.GetRecipes)
}
