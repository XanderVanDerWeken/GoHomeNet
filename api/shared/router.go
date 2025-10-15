package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/api/auth"
	"github.com/xandervanderweken/GoHomeNet/api/cards"
	"github.com/xandervanderweken/GoHomeNet/api/chores"
	"github.com/xandervanderweken/GoHomeNet/api/finances"
	"github.com/xandervanderweken/GoHomeNet/api/recipes"
	"github.com/xandervanderweken/GoHomeNet/api/users"
	"github.com/xandervanderweken/GoHomeNet/internal/container"
)

func NewRouter(c *container.Container) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/auth", func(r chi.Router) {
			auth.Routes(r, c.AuthSvc)
		})

		apiRouter.Group(func(protected chi.Router) {
			protected.Use(AuthMiddleware)

			protected.Route("/cards", func(r chi.Router) {
				cards.Routes(r, c.CardsSvc, c.UserSvc)
			})

			protected.Route("/chores", func(r chi.Router) {
				chores.Routes(r, c.ChoresSvc, c.UserSvc)
			})

			protected.Route("/finances", func(r chi.Router) {
				finances.Routes(r, c.FinancesSvc)
			})

			protected.Route("/recipes", func(r chi.Router) {
				recipes.Routes(r, c.RecipeSvc, c.UserSvc)
			})

			protected.Route("/users", func(r chi.Router) {
				users.Routes(r, c.UserSvc)
			})
		})
	})

	return router
}
