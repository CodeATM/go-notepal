package main

import (
	"log"

	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/database"
	"github.com/CodeATM/notepal-go/internal/middlewares"
	"github.com/CodeATM/notepal-go/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database.ConnectDb(cfg)
	// Initialize Fiber app
	app := fiber.New()

	// Add middlewares
	app.Use(logger.New())
	app.Use(middlewares.ErrorMiddleware())

	// Set up routes
	routes.SetupRoutes(app, cfg)

	// Start the server
	log.Println("Starting server on port :3000")
	log.Fatal(app.Listen(":3000"))
}
