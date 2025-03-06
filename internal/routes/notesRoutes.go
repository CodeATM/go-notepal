package routes

import (
	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/middlewares"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/CodeATM/notepal-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterNoteRoutes(api fiber.Router, cfg config.Config) {
	// Fetch all notes
	api.Get("/posts", func(c *fiber.Ctx) error {
		var posts []models.Note
		database.DB.Find(&posts)
		return c.JSON(posts)
	})

	// Create a new note with authentication
	api.Post("/note/create", middlewares.AuthMiddleware(cfg), func(c *fiber.Ctx) error {
		return service.CreateNote(c)
	})
}
