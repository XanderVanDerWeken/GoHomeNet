package chores

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type choreEventHandler struct {
	repo Repository
}

func NewChoreEventHandler(repo Repository) *choreEventHandler {
	return &choreEventHandler{
		repo: repo,
	}
}

func (h *choreEventHandler) handleChoreEvent(e events.Event) {
	switch event := e.(type) {
	case NewChoreEvent:
		h.repo.CreateChore(&event.NewChore)
	case CompletedChoreEvent:
		h.repo.CompleteChore(event.ChoreId)
	}
}
