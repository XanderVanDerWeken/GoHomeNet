package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/cards"
	"github.com/xandervanderweken/GoHomeNet/internal/chores"
	"github.com/xandervanderweken/GoHomeNet/internal/config"
	"github.com/xandervanderweken/GoHomeNet/internal/database"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func main() {
	config.Load("./config")

	db := database.Connect()
	db.AutoMigrate(&chores.Chore{}, &cards.Card{}, &users.User{})

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/chores", func(r chi.Router) {
			chores.RegisterRoutes(r, db)
		})

		r.Route("/cards", func(r chi.Router) {
			cards.RegisterRoutes(r, db)
		})
	})

	port := config.AppConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	log.Println("Server running in port", port)
	http.ListenAndServe(addr, r)
}
