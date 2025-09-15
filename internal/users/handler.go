package users

import (
	"encoding/json"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

type UserHandler struct {
	service Service
}

func NewUserHandler(service Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) PostSignupUser(w http.ResponseWriter, r *http.Request) {
	var dto SignupDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, shared.ErrBadRequest)
		return
	}

	h.service.SignUpUser(dto.Username, dto.Password, dto.FirstName, dto.LastName)
	w.WriteHeader(http.StatusCreated)
}
