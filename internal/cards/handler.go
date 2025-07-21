package cards

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	errx "github.com/xandervanderweken/GoHomeNet/internal/errors"
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
		errx.RespondError(w, err)
		return
	}

	json.NewEncoder(w).Encode(cards)
}

func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var req CreateCardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errx.RespondError(w, errx.ErrValidation)
		return
	}

	cardDto, err := h.service.CreateCard(req)
	if err != nil {
		errx.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cardDto)
}

func (h *CardHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	cardIdStr := chi.URLParam(r, "cardId")

	cardId64, err := strconv.ParseUint(cardIdStr, 10, 32)
	if err != nil {
		errx.RespondError(w, errx.ErrValidation)
		return
	}
	cardId := uint(cardId64)

	if err := h.service.DeleteCard(cardId); err != nil {
		errx.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
