package chores

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	choreService Service
}

func NewHandler(service Service) *Handler {
	return &Handler{choreService: service}
}

func (h *Handler) GetAllChores(w http.ResponseWriter, r *http.Request) {
	chores, err := h.choreService.GetAllChores()
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(chores)
}
