package utils

import (
	"encoding/json"
	"net/http"
)

// Simple type to recognize valid functions
type apiFunc func(w http.ResponseWriter, r *http.Request) error

// APIError Simple type for formatting errors
type APIError struct {
	Error string
}

// Creates simple handler with double error check (nobody wants its json fucked up)
func MakeHTTPHandleFunc(function apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := function(w, r); err != nil {
			// Global error handling
			err := WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
			if err != nil {
				http.Error(w, "Internal JSON writer error", http.StatusInternalServerError)
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, content any) error {
	// First set the type of header properly to json
	// Must be first otherwise the response goes gah-gah
	w.Header().Set("Content-Type", "application/json")

	// Write given status
	w.WriteHeader(status)

	// Return the whole thing
	// We must use NewEncoder in order to format information to JSON
	// then pass it the method .Encode(any given data)
	return json.NewEncoder(w).Encode(content)
}
