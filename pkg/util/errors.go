package util

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
}
