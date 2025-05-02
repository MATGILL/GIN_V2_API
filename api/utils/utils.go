package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, dto any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	err := json.NewDecoder(r.Body).Decode(dto)
	return err
}

func WriteJSON(w http.ResponseWriter, status int, output any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(output)

}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error:": err.Error()})
}
