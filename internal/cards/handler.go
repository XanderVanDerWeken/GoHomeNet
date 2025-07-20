package cards

import (
	"encoding/json"
	"net/http"
)

type CardHandler struct {
	service Service
}

func NewHandler(service Service) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) GetAllCards(w http.ResponseWriter, r *http.Request) {
	cards, err := h.service.GetAllCards()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cards)
}

func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var req CreateCardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cardDto, err := h.service.CreateCard(req)
	if err != nil {
		http.Error(w, "Failed to create card", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cardDto)
}
