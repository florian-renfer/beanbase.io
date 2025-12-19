package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")
	ErrInvalidInput     = errors.New("invalid input")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

// NewError creates a new Error instance with a single error message and HTTP status code.
func NewError(err error, status int) Error {
	return Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

// NewErrorMessage creates a new Error instance with multiple error messages and HTTP status code.
func NewErrorMessage(messages []string, status int) Error {
	return Error{
		statusCode: status,
		Errors:     messages,
	}
}

// Send writes the error response as JSON to the provided http.ResponseWriter.
func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
