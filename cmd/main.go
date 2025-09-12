package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/cards"
	"github.com/xandervanderweken/GoHomeNet/internal/chores"
	"github.com/xandervanderweken/GoHomeNet/internal/config"
	"github.com/xandervanderweken/GoHomeNet/internal/container"
	"github.com/xandervanderweken/GoHomeNet/internal/finances"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func main() {
	c := container.New()

	r := chi.NewRouter()

	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/cards", func(r chi.Router) {
			cards.Routes(r, c.CardsSvc, c.UserSvc)
		})

		apiRouter.Route("/chores", func(r chi.Router) {
			chores.Routes(r, c.ChoresSvc, c.UserSvc)
		})

		apiRouter.Route("/finances", func(r chi.Router) {
			finances.Routes(r, c.FinancesSvc)
		})

		apiRouter.Route("/users", func(r chi.Router) {
			users.Routes(r, c.UserSvc)
		})
	})

	port := config.AppConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	log.Println("Server running in port", port)
	http.ListenAndServe(addr, r)
}
