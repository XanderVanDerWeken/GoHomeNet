package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type UserHandler struct {
	service users.Service
}

func NewUserHandler(service users.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	dto := UserDto{
		Username:  user.Username,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		shared.WriteError(w, err)
		return
	}
}
