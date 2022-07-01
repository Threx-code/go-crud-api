package exceptions

import (
	"encoding/json"
	"net/http"
)

type ValidationError struct {
	Error string `json:"error"`
}

func NewValidationError(w http.ResponseWriter, errorCode int, errorMessage string) *ValidationError {
	validator := &ValidationError{
		Error: errorMessage,
	}

	if validator.Error != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorCode)
		json.NewEncoder(w).Encode(validator)
	}

	return validator
}
