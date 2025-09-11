package finances

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

func Routes(service Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/categories", func(w http.ResponseWriter, r *http.Request) {
		var dto struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			writeError(w, err)
			return
		}

		if err := service.CreateCategory(dto.Name); err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		categories := service.GetAllCategories()

		categoryDtos := make([]CategoryDto, 0, len(categories))

		for _, category := range categories {
			categoryDtos = append(categoryDtos, CategoryDto{
				Name: category.Name,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(categoryDtos); err != nil {
			writeError(w, err)
			return
		}
	})

	return r
}

func writeError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*shared.AppError); ok {
		http.Error(w, appErr.Message, appErr.Status)
		return
	}
	http.Error(w, shared.ErrInternal.Message, shared.ErrInternal.Status)
}
