package util

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

var ErrNotFound = errors.New("not found")

func (e *AppError) Error() string {
	return e.Message
}

func NotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func UnauthorizedError(message string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: message}
}

func BadRequestError(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

func InternalServerError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func ConflictError(message string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: message}
}
