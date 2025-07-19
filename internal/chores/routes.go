package chores

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(r chi.Router, db *gorm.DB) {
	repo := NewChoreRepository(db)
	service := NewChoreService(repo)
	handler := NewChoreHandler(service)

	r.Get("/chores", handler.GetAllChores)
}
