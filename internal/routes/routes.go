package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/CodeATM/notepal-go/config"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("This server is healthy")
}

func SetupRoutes(app *fiber.App, cfg config.Config) {
	api := app.Group("/api/v1")

	// Register all routes
	api.Get("/healthz", welcome)
	RegisterUserRoutes(api)
	RegisterNoteRoutes(api)
	RegisterAuthRoutes(api, cfg)
}
