package chores

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

func Routes(service Service, userService users.Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var dto struct {
			Username string    `json:"username"`
			Title    string    `json:"title"`
			Notes    string    `json:"notes"`
			DueDate  time.Time `json:"dueDate"`
		}
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			writeError(w, err)
			return
		}

		service.CreateChore(dto.Username, dto.Title, dto.Notes, &dto.DueDate)
	})

	return r
}

func writeError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*shared.AppError); ok {
		http.Error(w, appErr.Message, appErr.Status)
		return
	}
	http.Error(w, shared.ErrInternal.Message, shared.ErrInternal.Status)
}
