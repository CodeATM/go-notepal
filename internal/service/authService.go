package service

import (
	"time"

	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/models"
	"github.com/CodeATM/notepal-go/internal/utils"
	"github.com/CodeATM/notepal-go/internal/utils/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// RegisterService handles user registration
func RegisterService(c *fiber.Ctx, cfg config.Config) error {
	// Parse request body into RegisterRequest struct
	req := new(requests.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" || req.Firstname == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Email, password, and firstname are required", nil)
	}

	// Check for unique email
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, "Email already exists", nil)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}

	// Create new user object
	user := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	// Save the user
	if result := database.DB.Create(&user); result.Error != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil)
	}

	// Generate JWT token
	token, err := generateJWT(user.ID.String(), cfg.JwtSecret)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	// Response data
	data := fiber.Map{
		"id":    user.ID.String(),
		"token": token,
		"name":  user.Firstname,
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", data)
}

func LoginService(c *fiber.Ctx, cfg config.Config) error {
	// Parse request body
	req := new(requests.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Email and password are required", nil)
	}

	// Find user
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid password", nil)
	}

	// Generate JWT token
	token, err := generateJWT(user.ID.String(), cfg.JwtSecret)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	// Success response
	data := fiber.Map{
		"id":    user.ID.String(),
		"token": token,
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "User logged in successfully", data)
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
