package util

import (
	"encoding/json"
	"net/http"

	"github.com/krishnapramodaradhi/event-booking-api/internal/config"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request, db *config.Database) (int, error)

func MakeHttpHandleFunc(f handlerFunc, db *config.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if statusCode, err := f(w, r, db); err != nil {
			WriteJSON(w, statusCode, map[string]any{"error": err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
