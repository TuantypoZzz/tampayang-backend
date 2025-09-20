package middleware

import (
	"fmt"

	"tampayang-backend/core/errors"
	"tampayang-backend/core/response"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a middleware that handles all errors in the application
func ErrorHandler(c *fiber.Ctx) error {
	// Catch any panic and convert it to a proper error
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic recovered: %v", r)
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Internal Server Error",
				"message": "An unexpected error occurred",
				"details": err.Error(),
			})
		}
	}()

	// Continue with the next middleware
	return c.Next()
}

// HandleError is a utility function to handle errors consistently
func HandleError(c *fiber.Ctx, err error) error {
	// If it's already our custom error type, use it directly
	if appErr, ok := err.(*errors.AppError); ok {
		return response.ErrorResponse(c, appErr)
	}

	// For other errors, wrap them as internal errors
	appErr := errors.Wrap(err, "INTERNAL_ERROR", "An unexpected error occurred", fiber.StatusInternalServerError)
	return response.ErrorResponse(c, appErr)
}