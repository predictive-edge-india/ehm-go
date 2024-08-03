package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	authHandlers "github.com/predictive-edge-india/ehm-go/handlers/auth"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/signin", func(c *fiber.Ctx) error {
		return authHandlers.SigninWithPassword(c)
	})
	auth.Post("/signup", func(c *fiber.Ctx) error {
		return authHandlers.SignupUser(c)
	})
	auth.Get("/validate", middlewares.Protected(), authHandlers.ValidateUser)
}
