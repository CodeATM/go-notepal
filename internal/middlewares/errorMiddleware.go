package middlewares

import (
	"errors"

	"github.com/CodeATM/notepal-go/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	// Replace with your actual project path
)

func ErrorMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if rec := recover(); rec != nil {
				err := rec.(error)
				handleError(c, err)
			}
		}()
		return c.Next()
	}
}

func handleError(c *fiber.Ctx, err error) error {
	// Handle application-specific errors
	if appErr, ok := err.(*utils.AppError); ok {
		return utils.ErrorResponse(c, appErr.StatusCode, appErr.Message, nil)
	}

	// Handle Postgres errors
	var pqError *pq.Error
	if errors.As(err, &pqError) {
		switch pqError.Code {
		case "23505": // Unique violation
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Duplicate entry detected.", nil)
		case "23514": // Check violation
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Constraint violation occurred.", nil)
		default:
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Database error.", nil)
		}
	}

	// For unhandled errors, return an internal server error
	return utils.ErrorResponse(c, fiber.StatusInternalServerError, "An unexpected error occurred.", nil)
}
