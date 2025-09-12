package cards

import (
	"encoding/json"
	"net/http"
	"time"

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
	var dto struct {
		Username  string    `json:"username"`
		Name      string    `json:"name"`
		ExpiresAt time.Time `json:"expiresAt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeError(w, shared.ErrBadRequest)
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
		writeError(w, err)
		return
	}

	cardDtos := make([]CardDto, 0, len(cards))

	for _, card := range cards {
		uName := username
		if uName == "" {
			user, err := h.userService.GetUserByUserId(card.UserID)
			if err != nil {
				writeError(w, err)
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
		writeError(w, err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*shared.AppError); ok {
		http.Error(w, appErr.Message, appErr.Status)
		return
	}
	http.Error(w, shared.ErrInternal.Message, shared.ErrInternal.Status)
}
