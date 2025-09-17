package users

import (
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

var (
	ErrUserNotFound = shared.NewAppError("USER_NOT_FOUND", "user not found", http.StatusNotFound)
)
