package auth

import (
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

var (
	ErrMissingHeader     = shared.NewAppError("AUTH_HEADER_MISSING", "Missing Authorization header", http.StatusUnauthorized)
	ErrWrongHeaderFormat = shared.NewAppError("AUTH_WRONG_FORMAT", "Invalid Authorization header format", http.StatusUnauthorized)
	ErrInvalidToken      = shared.NewAppError("AUTH_INVALID_TOKEN", "Invalid token", http.StatusUnauthorized)
)
