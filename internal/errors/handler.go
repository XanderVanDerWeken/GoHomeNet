package errors

import (
	"encoding/json"
	"net/http"
)

func RespondError(w http.ResponseWriter, err error) {
	dto := ToDTO(err)
	status := StatusCode(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto)
}
