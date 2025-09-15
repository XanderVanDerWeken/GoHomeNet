package shared

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteError(w http.ResponseWriter, err error) {
	var appError *AppError

	if errors.As(err, &appError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(appError.Status)
		json.NewEncoder(w).Encode(appError)
	} else {
		http.Error(w, ErrInternal.Message, ErrInternal.Status)
	}
}
