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
	"github.com/xandervanderweken/GoHomeNet/internal/finances"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func main() {
	config.Load("./config")

	// Connect to the database
	db := database.Connect()
	db.AutoMigrate(&users.User{}, &cards.Card{}, &chores.Chore{}, &finances.Transaction{}, &finances.Category{})

	// Add Users Module
	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)

	// Add Cards Module
	cardRepo := cards.NewRepository(db)
	cardService := cards.NewService(cardRepo, userRepo)

	// Add Chores Module
	choreRepo := chores.NewRepository(db)
	choreService := chores.NewService(choreRepo, userRepo)

	// Add Finances Module
	transactionRepo := finances.NewTransactionRepository(db)
	categoryRepo := finances.NewCategoryRepository(db)
	financesService := finances.NewService(transactionRepo, categoryRepo)

	r := chi.NewRouter()

	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/cards", func(r chi.Router) {
			cards.Routes(r, cardService, userService)
		})

		apiRouter.Route("/chores", func(r chi.Router) {
			chores.Routes(r, choreService, userService)
		})

		apiRouter.Route("/finances", func(r chi.Router) {
			finances.Routes(r, financesService)
		})

		apiRouter.Route("/users", func(r chi.Router) {
			users.Routes(r, userService)
		})
	})

	port := config.AppConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	log.Println("Server running in port", port)
	http.ListenAndServe(addr, r)
}
