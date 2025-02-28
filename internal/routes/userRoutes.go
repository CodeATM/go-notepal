package routes

import (
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/CodeATM/notepal-go/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router) {
	api.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(users)
	})

	api.Post("/users", func(c *fiber.Ctx) error {
		user := new(models.User)

		// Parse the request body
		if err := c.BodyParser(user); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, utils.BadRequestError("Invalid request body").Error(), nil)
		}

		// Validate required fields
		if user.Email == "" {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, utils.BadRequestError("Email is required").Error(), nil)
		}

		if len(user.Password) < 6 {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, utils.BadRequestError("Password must be at least 6 characters long").Error(), nil)
		}

		// Check for unique email
		var existingUser models.User
		if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			return utils.ErrorResponse(c, fiber.StatusConflict, utils.BadRequestError("Email already exists").Error(), nil)
		}

		// Save the user
		if result := database.DB.Create(user); result.Error != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.InternalServerError("Failed to create user").Error(), nil)
		}

		// Return success response
		return utils.SuccessResponse(c, fiber.StatusCreated, "User created successfully", user)
	})
}
