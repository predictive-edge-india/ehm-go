package routes

import (
	"github.com/gofiber/fiber/v2"
	dashboardHandlers "github.com/predictive-edge-india/ehm-go/handlers/dashboard"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func DashboardRoutes(app fiber.Router) {
	group := app.Group("/dashboard")
	group.Get("/home", middlewares.Protected(), dashboardHandlers.FetchDashboardHome)
}
