package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

func Routes(service Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
		var dto struct {
			Username  string `json:"username"`
			Password  string `json:"password"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		}
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			writeError(w, shared.ErrBadRequest)
			return
		}

		service.SignUpUser(dto.Username, dto.Password, dto.FirstName, dto.LastName)
		w.WriteHeader(http.StatusCreated)
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
