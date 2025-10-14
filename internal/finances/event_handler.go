package finances

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type CategoryEventHandler struct {
	repo CategoryRepository
}

func (h *CategoryEventHandler) HandleNewCategory(event events.Event) {
	if event, ok := event.(NewCategoryEvent); ok {
		h.repo.AddCategory(event.NewName)
	}
}
