package errors

import (
	"errors"
	"net/http"
)

func StatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case "NOT_FOUND":
			return http.StatusNotFound
		case "UNAUTHORIZED":
			return http.StatusUnauthorized
		case "VALIDATION_ERROR":
			return http.StatusBadRequest
		case "INTERNAL_SERVER_ERROR":
			return http.StatusInternalServerError
		default:
			return http.StatusInternalServerError
		}
	}

	return http.StatusInternalServerError
}

func ToDTO(err error) ErrorResponse {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Details: unwrapMessage(appErr.Err),
		}
	}

	return ErrorResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "An unexpected error occurred",
	}
}

func unwrapMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
