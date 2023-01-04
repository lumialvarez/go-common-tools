package apierrors

import "net/http"

func NewConflictError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusConflict,
		Message: message,
		Err:     "Conflict Error",
		Cause:   cause,
	}
}
