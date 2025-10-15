package finances

import (
	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/finances"
)

func Routes(router chi.Router, service finances.Service) {
	handler := NewFinanceHandler(service)

	router.Post("/categories", handler.PostNewCategory)
	router.Get("/categories", handler.GetAllCategories)
}
