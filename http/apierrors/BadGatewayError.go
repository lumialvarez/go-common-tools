package apierrors

import "net/http"

func NewBadGatewayError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusBadGateway,
		Message: message,
		Err:     "Bad Gateway Error",
		Cause:   cause,
	}
}
