package errors

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const (
	InternalServerError = "Internal Server Error"
	NotFound            = "Not Found"
	Unauthorized        = "Unauthorized"
)

// Error implements error interface
// so that we can use Error struct as error type
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// Status returns the status code of the error
func (e *Error) Status() int {
	return e.Code
}

func NewError(message string, code int) *Error {
	return &Error{Message: message, Code: code}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInternalServerError() *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: InternalServerError,
	}
}

func NewNotFoundError() *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: NotFound,
	}
}

func NewUnauthorizedError() *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: Unauthorized,
	}
}
