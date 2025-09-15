package finances

import (
	"encoding/json"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

type FinanceHandler struct {
	service Service
}

func NewFinanceHandler(service Service) *FinanceHandler {
	return &FinanceHandler{service: service}
}

func (h *FinanceHandler) PostNewCategory(w http.ResponseWriter, r *http.Request) {
	var dto NewCategoryDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, err)
		return
	}

	newCategory := Category{
		Name: dto.Name,
	}
	if err := h.service.CreateCategory(&newCategory); err != nil {
		shared.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *FinanceHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.service.GetAllCategories()

	categoryDtos := make([]CategoryDto, 0, len(categories))

	for _, category := range categories {
		categoryDtos = append(categoryDtos, CategoryDto{
			Name: category.Name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(categoryDtos); err != nil {
		shared.WriteError(w, err)
		return
	}
}
