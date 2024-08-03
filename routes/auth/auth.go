package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	authHandlers "github.com/predictive-edge-india/ehm-go/handlers/auth"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AuthRoutes(router fiber.Router) {
	group := router.Group("/auth")
	group.Post("/signin", func(c *fiber.Ctx) error {
		return authHandlers.SigninWithPassword(c)
	})
	group.Post("/signup", func(c *fiber.Ctx) error {
		return authHandlers.SignupUser(c)
	})
	group.Get("/validate", middlewares.Protected(), authHandlers.ValidateUser)
}
