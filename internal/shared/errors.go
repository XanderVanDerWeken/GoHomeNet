package shared

import "net/http"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

var (
	ErrNotFound     = NewAppError("NOT_FOUND", "resource not found", http.StatusNotFound)
	ErrBadRequest   = NewAppError("BAD_REQUEST", "invalid request", http.StatusBadRequest)
	ErrInternal     = NewAppError("INTERNAL", "internal server error", http.StatusInternalServerError)
	ErrUnauthorized = NewAppError("UNAUTHORIZED", "Unauthorized request", http.StatusUnauthorized)
)
