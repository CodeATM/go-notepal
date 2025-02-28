package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/CodeATM/notepal-go/config"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	api := app.Group("/api/v1")

	// Register all routes
	RegisterUserRoutes(api)
	RegisterNoteRoutes(api)
	RegisterAuthRoutes(api, cfg)
}
