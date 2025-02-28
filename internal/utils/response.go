package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// SuccessResponse formats a successful HTTP response in Fiber.
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	if statusCode < 200 || statusCode > 299 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid status code. Use a valid status code.",
			"error":   true,
		})
	}

	if !strings.HasSuffix(message, ".") {
		message += "."
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"data":    data,
		"error":   false,
	})
}

// ErrorResponse formats an error HTTP response in Fiber.
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	if !strings.HasSuffix(message, ".") {
		message += "."
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message":    message,
		"statusCode": statusCode,
		"data":       data,
		"error":      true,
	})
}
