package vndb

import (
	"fmt"
	"net/http"
)

type Error struct {
	apiError *apiError
	msg      string
}

func (e Error) Is(apiError apiError) bool {
	return e.apiError != nil && e.apiError.StatusCode == apiError.StatusCode
}

func (e Error) Error() string {
	if e.apiError != nil {
		return fmt.Sprintf("%s: %s", e.apiError.Error(), e.msg)
	}
	return e.msg
}

type apiError struct {
	StatusCode int
	Reason     string
}

func (e apiError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Reason)
}

var (
	ErrBadRequest = apiError{
		StatusCode: http.StatusBadRequest,
		Reason:     "Invalid request body or query, the included error message hopefully points at the problem.",
	}
	ErrUnauthorized = apiError{
		StatusCode: http.StatusUnauthorized,
		Reason:     "Invalid authentication token.",
	}
	ErrNotFound = apiError{
		StatusCode: http.StatusNotFound,
		Reason:     "Invalid API path or HTTP method.",
	}
	ErrTooManyRequests = apiError{
		StatusCode: http.StatusTooManyRequests,
		Reason:     "Throttled.",
	}
	ErrInternalServerError = apiError{
		StatusCode: http.StatusInternalServerError,
		Reason:     "Server error, usually points to a bug if this persists.",
	}
	ErrBadGateway = apiError{
		StatusCode: http.StatusBadGateway,
		Reason:     "Server is down, should be temporary.",
	}
)
