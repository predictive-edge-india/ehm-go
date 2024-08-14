package routes

import (
	"github.com/gofiber/fiber/v2"
	userHandlers "github.com/predictive-edge-india/ehm-go/handlers/users"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func UserRoutes(app fiber.Router) {
	users := app.Group("/users")
	users.Get("/formdata", middlewares.Protected(), userHandlers.FetchUserFormData)
	users.Get("/:userId", middlewares.Protected(), userHandlers.FetchUserDetails)
	users.Delete("/:userId", middlewares.Protected(), userHandlers.DeleteUser)
	users.Get("/", middlewares.Protected(), userHandlers.FetchUsers)
	users.Post("/", middlewares.Protected(), userHandlers.CreateNewUser)
}
