package chores

import (
	errx "github.com/xandervanderweken/GoHomeNet/internal/errors"

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chores)
}

func (h *Handler) CreateChore(w http.ResponseWriter, r *http.Request) {
	var req CreateChoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errx.RespondError(w, errx.ErrValidation)
		return
	}

	choresDto, err := h.choreService.CreateChore(req)
	if err != nil {
		errx.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(choresDto)
}
