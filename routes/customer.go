package routes

import (
	"github.com/gofiber/fiber/v2"
	customerHandlers "github.com/predictive-edge-india/ehm-go/handlers/customers"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func CustomerRoutes(app fiber.Router) {
	group := app.Group("/customers")
	group.Get("/:customerId/stats", middlewares.Protected(), customerHandlers.FetchCustomerStats)
	group.Get("/:customerId", middlewares.Protected(), customerHandlers.FetchCustomerDetails)
	group.Patch("/:customerId", middlewares.Protected(), customerHandlers.UpdateCustomerDetails)
	group.Get("/", middlewares.Protected(), customerHandlers.FetchAllCustomers)
	group.Post("/", middlewares.Protected(), customerHandlers.CreateNewCustomer)
}
