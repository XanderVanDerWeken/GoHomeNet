package recipes

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type recipeEventHandler struct {
	repo Repository
}

func NewRecipeEventHandler(repo Repository) *recipeEventHandler {
	return &recipeEventHandler{repo: repo}
}

func (h *recipeEventHandler) handleRecipeEvent(e events.Event) {
	if event, ok := e.(NewRecipeEvent); ok {
		h.repo.CreateRecipe(&event.Recipe)
	}
}
