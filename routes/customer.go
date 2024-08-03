package routes

import (
	"github.com/gofiber/fiber/v2"
	customerHandlers "github.com/predictive-edge-india/ehm-go/handlers/customers"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func CustomerRoutes(app fiber.Router) {
	customers := app.Group("/customers")
	customers.Get("/", middlewares.Protected(), customerHandlers.FetchAllCustomers)
	customers.Post("/", middlewares.Protected(), customerHandlers.CreateNewCustomer)
}
