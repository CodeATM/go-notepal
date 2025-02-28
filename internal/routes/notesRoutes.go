package routes

import (
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterNoteRoutes(api fiber.Router) {
	api.Get("/posts", func(c *fiber.Ctx) error {
		var posts []models.Note
		database.DB.Find(&posts)
		return c.JSON(posts)
	})

	api.Post("/posts", func(c *fiber.Ctx) error {
		post := new(models.Note)

		// Parse the request body
		if err := c.BodyParser(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		// Save the post
		if result := database.DB.Create(post); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create post",
			})
		}

		return c.JSON(post)
	})
}
