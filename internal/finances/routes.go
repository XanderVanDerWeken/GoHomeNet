package finances

import "github.com/go-chi/chi/v5"

func Routes(router chi.Router, service Service) {
	handler := NewFinanceHandler(service)

	router.Post("/categories", handler.PostNewCategory)
	router.Get("/categories", handler.GetAllCategories)
}
