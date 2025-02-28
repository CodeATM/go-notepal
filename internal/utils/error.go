package utils

import "net/http"

// AppError defines the structure for application-specific errors.
type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Helper functions to create specific errors
func UnauthorizedError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func BadRequestError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func ForbiddenError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func InternalServerError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}
