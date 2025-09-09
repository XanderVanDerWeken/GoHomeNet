package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/config"
	"github.com/xandervanderweken/GoHomeNet/internal/database"
)

func main() {
	config.Load("./config")

	db := database.Connect()
	db.AutoMigrate()

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/chores", func(r chi.Router) {
			//chores.RegisterRoutes(r, db)
		})

		r.Route("/cards", func(r chi.Router) {
			//cards.RegisterRoutes(r, db)
		})
	})

	port := config.AppConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	log.Println("Server running in port", port)
	http.ListenAndServe(addr, r)
}
