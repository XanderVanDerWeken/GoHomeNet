package chores

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(r chi.Router, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	r.Get("/", handler.GetAllChores)
	r.Post("/", handler.CreateChore)
}
