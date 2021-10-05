package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPError ...
type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}

// Execute returns HTTP error response
func Execute(w http.ResponseWriter, data interface{}, status int, err error) {
	if err != nil {
		data = HTTPError{err.Error()}
	}
	if status != http.StatusFound {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("could not encode response to output: %v", err)
	}
}
