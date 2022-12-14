package apierrors

import "net/http"

func NewForbiddenError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusForbidden,
		Message: message,
		Err:     "Forbidden Error",
		Cause:   cause,
	}
}
