package auth

import (
	"encoding/json"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/auth"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *AuthHandler {
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

	token, err := h.service.SignUpUser(newUser)
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(AuthDto{Token: *token}); err != nil {
		shared.WriteError(w, err)
		return
	}
}

func (h *AuthHandler) PostLoginUser(w http.ResponseWriter, r *http.Request) {
	var dto LoginDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, shared.ErrBadRequest)
		return
	}

	token, err := h.service.LoginUser(dto.Username, dto.Password)
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(AuthDto{Token: *token}); err != nil {
		shared.WriteError(w, err)
		return
	}
}
