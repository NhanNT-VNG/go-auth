package helpers

import "net/http"

type APIError interface {
	APIError() (int, string)
}

type sentinelAPIError struct {
	status  int
	message string
}

type sentinelWrappedError struct {
	error
	sentinel *sentinelAPIError
}

var (
	ErrAuth       = &sentinelAPIError{status: http.StatusUnauthorized, message: "unauthorize"}
	ErrForbidden  = &sentinelAPIError{status: http.StatusForbidden, message: "forbidden"}
	ErrNotFound   = &sentinelAPIError{status: http.StatusNotFound, message: "not found"}
	ErrBadRequest = &sentinelAPIError{status: http.StatusBadRequest, message: "bad request"}
)
