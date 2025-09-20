package errors

import (
	"fmt"
	"net/http"
)

// AppError represents a custom application error
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// New creates a new AppError
func New(code, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// NewWithDetails creates a new AppError with details
func NewWithDetails(code, message, details string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
		Status:  status,
	}
}

// Wrap creates a new AppError that wraps another error
func Wrap(err error, code, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: err.Error(),
		Status:  status,
	}
}

// Common error types
var (
	ErrBadRequest          = New("BAD_REQUEST", "Bad request", http.StatusBadRequest)
	ErrUnauthorized        = New("UNAUTHORIZED", "Unauthorized", http.StatusUnauthorized)
	ErrForbidden           = New("FORBIDDEN", "Forbidden", http.StatusForbidden)
	ErrNotFound            = New("NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrInternalServer      = New("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
	ErrServiceUnavailable  = New("SERVICE_UNAVAILABLE", "Service unavailable", http.StatusServiceUnavailable)
	ErrValidationError     = New("VALIDATION_ERROR", "Validation failed", http.StatusBadRequest)
	ErrDatabaseError       = New("DATABASE_ERROR", "Database operation failed", http.StatusInternalServerError)
	ErrFileUploadError     = New("FILE_UPLOAD_ERROR", "File upload failed", http.StatusInternalServerError)
	ErrNotificationError   = New("NOTIFICATION_ERROR", "Notification sending failed", http.StatusInternalServerError)
)