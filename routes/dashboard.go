package routes

import (
	"github.com/gofiber/fiber/v2"
	dashboardHandlers "github.com/predictive-edge-india/ehm-go/handlers/dashboard"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func DashboardRoutes(app fiber.Router) {
	group := app.Group("/dashboard")
	group.Get("/asset-locations", middlewares.Protected(), dashboardHandlers.FetchAssetLocations)
	group.Get("/customer-locations", middlewares.Protected(), dashboardHandlers.FetchCustomerLocations)
	group.Get("/home", middlewares.Protected(), dashboardHandlers.FetchDashboardHome)
}
