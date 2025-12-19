package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     any
}

// NewSuccess creates a new Success instance with the given result and HTTP status code.
func NewSuccess(result any, status int) Success {
	return Success{
		statusCode: status,
		result:     result,
	}
}

// Send writes the success response as JSON to the provided http.ResponseWriter.
func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	return json.NewEncoder(w).Encode(r.result)
}
