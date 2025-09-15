package cards

import (
	"encoding/json"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type CardHandler struct {
	service     Service
	userService users.Service
}

func NewCardHandler(service Service, userService users.Service) *CardHandler {
	return &CardHandler{service: service, userService: userService}
}

func (h *CardHandler) PostNewCard(w http.ResponseWriter, r *http.Request) {
	var dto NewCardDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, shared.ErrBadRequest)
		return
	}

	h.service.AddCard(dto.Username, dto.Name, dto.ExpiresAt)
}

func (h *CardHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	username := q.Get("username")

	var cards []Card
	var err error

	if username != "" {
		cards, err = h.service.GetAllOwnCards(username)
	} else {
		cards = h.service.GetAllCards()
	}
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	cardDtos := make([]CardDto, 0, len(cards))

	for _, card := range cards {
		uName := username
		if uName == "" {
			user, err := h.userService.GetUserByUserId(card.UserID)
			if err != nil {
				shared.WriteError(w, err)
				return
			}
			uName = user.Username
		}

		cardDtos = append(cardDtos, CardDto{
			Username:  uName,
			Name:      card.Name,
			ExpiresAt: card.ExpiresAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cardDtos); err != nil {
		shared.WriteError(w, err)
		return
	}
}
