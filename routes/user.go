package routes

import (
	"github.com/gofiber/fiber/v2"
	userHandlers "github.com/predictive-edge-india/ehm-go/handlers/users"
)

func UserRoutes(app fiber.Router) {
	users := app.Group("/users")
	users.Get("/", userHandlers.FetchUsers)
}
