package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/auth"
	"github.com/xandervanderweken/GoHomeNet/internal/cards"
	"github.com/xandervanderweken/GoHomeNet/internal/chores"
	"github.com/xandervanderweken/GoHomeNet/internal/container"
	"github.com/xandervanderweken/GoHomeNet/internal/finances"
	"github.com/xandervanderweken/GoHomeNet/internal/recipes"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func NewRouter(c *container.Container) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/cards", func(r chi.Router) {
			cards.Routes(r, c.CardsSvc, c.UserSvc)
		})

		apiRouter.Route("/chores", func(r chi.Router) {
			chores.Routes(r, c.ChoresSvc, c.UserSvc)
		})

		apiRouter.Route("/finances", func(r chi.Router) {
			finances.Routes(r, c.FinancesSvc)
		})

		apiRouter.Route("/recipes", func(r chi.Router) {
			recipes.Routes(r, c.RecipeSvc, c.UserSvc)
		})

		apiRouter.Route("/auth", func(r chi.Router) {
			auth.Routes(r, c.UserSvc)
		})

		apiRouter.Route("/users", func(r chi.Router) {
			users.Routes(r, c.UserSvc)
		})
	})

	return router
}
