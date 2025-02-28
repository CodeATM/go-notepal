package service

import (
	"time"

	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/CodeATM/notepal-go/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// RegisterService handles user registration
func RegisterService(c *fiber.Ctx, cfg config.Config) error {
	// Parse request body
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Validate required fields
	if user.Email == "" || user.Password == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Email and password are required", nil)
	}

	// Check for unique email
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, "Email already exists", nil)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}
	user.Password = string(hashedPassword)

	// Save the user
	if result := database.DB.Create(user); result.Error != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil)
	}

	// Generate JWT token using the user's UUID
	token, err := generateJWT(user.ID.String(), cfg.JwtSecret)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}
	data := fiber.Map{
		"id":    user.ID.String(), // Convert UUID to string
		"token": token,
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", data)
}

// Helper function to generate JWT token
func generateJWT(userID string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,                                // Use the user's UUID as a claim
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
