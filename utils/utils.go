package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJson(r *http.Request, dto any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err == io.EOF {
		return fmt.Errorf("empty request body")
	}
	if err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, output any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(output)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if jsonErr := WriteJSON(w, status, map[string]string{"error": err.Error()}); jsonErr != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
