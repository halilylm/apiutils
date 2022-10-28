package errors_test

import (
	"github.com/halilylm/apiutils/http/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("testing properties", func(t *testing.T) {
		msg := "bad gateway"
		statusBadGateway := errors.NewError(msg, http.StatusBadGateway)
		assert.EqualValues(t, http.StatusBadGateway, statusBadGateway.Status())
		assert.EqualValues(t, msg, statusBadGateway.Message)
	})
	t.Run("test if type implements the error interface", func(t *testing.T) {
		msg := "bad gateway"
		statusBadGateway := errors.NewError(msg, http.StatusBadGateway)
		var response any = statusBadGateway
		_, ok := response.(error)
		assert.EqualValues(t, ok, true)
	})
	t.Run("test error returns proper json", func(t *testing.T) {
		msg := "bad gateway"
		statusBadGateway := errors.NewError(msg, http.StatusBadGateway)
		expected := `{"message": "bad gateway", "code": 502}`
		actual := statusBadGateway.Error()
		assert.JSONEq(t, expected, actual)
	})
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Internal Server Error", errors.InternalServerError)
	assert.EqualValues(t, "Not Found", errors.NotFound)
	assert.EqualValues(t, "Unauthorized", errors.Unauthorized)
}

func TestNewBadRequestError(t *testing.T) {
	message := "bad request error"
	err := errors.NewBadRequestError(message)
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
}

func TestNewInternalServerError(t *testing.T) {
	message := errors.InternalServerError
	err := errors.NewInternalServerError()
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
}

func TestNewNotFoundError(t *testing.T) {
	message := errors.NotFound
	err := errors.NewNotFoundError()
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.Code)
}

func TestNewUnauthorizedError(t *testing.T) {
	message := errors.Unauthorized
	err := errors.NewUnauthorizedError()
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusUnauthorized, err.Code)
}
