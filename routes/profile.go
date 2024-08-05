package routes

import (
	"github.com/gofiber/fiber/v2"
	profileHandlers "github.com/predictive-edge-india/ehm-go/handlers/profile"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func ProfileRoutes(app fiber.Router) {
	group := app.Group("/profile")
	group.Get("/customers", middlewares.Protected(), func(c *fiber.Ctx) error {
		return profileHandlers.FetchProfileCustomer(c)
	})
}
