package apierrors

import "net/http"

func NewInternalServerError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Err:     "Internal Server Error",
		Cause:   cause,
	}
}
