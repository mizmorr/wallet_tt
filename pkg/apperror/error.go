package apperror

import "net/http"

type AppError struct {
	InternalCode int    `json:"code"`
	StatusCode   int    `json:"status_code"`
	Message      string `json:"message"`
	Details      string `json:"details,omitempty"`
}

var (
	ErrBadRequest = AppError{InternalCode: 1001, StatusCode: http.StatusBadRequest, Message: "Invalid input"}
	ErrValidation = AppError{InternalCode: 1002, StatusCode: http.StatusUnprocessableEntity, Message: "Validation error"}
	ErrForbidden  = AppError{InternalCode: 1003, StatusCode: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound   = AppError{InternalCode: 1005, StatusCode: http.StatusNotFound, Message: "Resource not found"}
)
