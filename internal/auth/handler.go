package auth

import (
	"encoding/json"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type AuthHandler struct {
	service Service
}

func NewAuthHandler(service Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) PostSignupUser(w http.ResponseWriter, r *http.Request) {
	var dto SignupDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, shared.ErrBadRequest)
		return
	}

	newUser := &users.User{
		Username:  dto.Username,
		Password:  dto.Password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	h.service.SignUpUser(newUser)
	w.WriteHeader(http.StatusCreated)
}
