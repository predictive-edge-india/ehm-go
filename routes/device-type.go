package routes

import (
	"github.com/gofiber/fiber/v2"
	deviceTypeHandlers "github.com/predictive-edge-india/ehm-go/handlers/deviceTypes"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func DeviceTypeRoutes(app fiber.Router) {
	group := app.Group("/device-types")
	group.Delete("/:deviceTypeId", middlewares.Protected(), deviceTypeHandlers.DeleteDeviceType)
	group.Get("/", middlewares.Protected(), deviceTypeHandlers.FetchDeviceTypes)
	group.Post("/", middlewares.Protected(), deviceTypeHandlers.CreateNewDeviceType)
}
