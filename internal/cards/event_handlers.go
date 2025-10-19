package cards

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type cardEventHandler struct {
	repo Repository
}

func NewCardEventHandler(repo Repository) *cardEventHandler {
	return &cardEventHandler{
		repo: repo,
	}
}

func (h *cardEventHandler) handleCardEvent(e events.Event) {
	if event, ok := e.(NewCardEvent); ok {
		h.repo.AddCard(&event.NewCard)
	}
}
