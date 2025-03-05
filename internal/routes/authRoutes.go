package routes

import (
	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(api fiber.Router, cfg config.Config) {
	api.Post("/auth/register", func(c *fiber.Ctx) error {
		return service.RegisterService(c, cfg)
	})
	api.Post("/auth/login", func(c *fiber.Ctx) error {
		return service.LoginService(c, cfg)
	})
}
