package service

import (
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/CodeATM/notepal-go/internal/utils"
	"github.com/CodeATM/notepal-go/internal/utils/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateNote(c *fiber.Ctx) error {
	// Retrieve user ID from context
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User ID not found or invalid", nil)
	}

	// Convert user ID from string to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid user ID format", nil)
	}

	// Parse request body
	req := new(requests.CreateNoteRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Validate request data
	if req.Content == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Ensure that the body contains content", nil)
	}

	// Create and save note
	note := models.Note{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := database.DB.Create(&note).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create note", err)
	}

	// Return response **without the User object
	return utils.SuccessResponse(c, fiber.StatusCreated, "Note added successfully", fiber.Map{
		"id":      note.ID,
		"title":   note.Title,
		"content": note.Content,
		"user_id": note.UserID, // Only include user_id, not user details
	})
}
