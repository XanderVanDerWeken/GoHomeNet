package chores

import (
	"encoding/json"
	"net/http"
)

type ChoreHandler struct {
	choreService ChoreService
}

func NewChoreHandler(service ChoreService) *ChoreHandler {
	return &ChoreHandler{choreService: service}
}

func (h *ChoreHandler) GetAllChores(w http.ResponseWriter, r *http.Request) {
	chores, err := h.choreService.GetAllChores()
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(chores)
}
