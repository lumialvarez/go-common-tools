package apierrors

import "net/http"

func NewNotFoundError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusNotFound,
		Message: message,
		Err:     "Bad Request Error",
		Cause:   cause,
	}
}
