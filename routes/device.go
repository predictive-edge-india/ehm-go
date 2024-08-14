package routes

import (
	"github.com/gofiber/fiber/v2"
	deviceHandlers "github.com/predictive-edge-india/ehm-go/handlers/devices"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func DeviceRoutes(app fiber.Router) {
	group := app.Group("/devices")
	group.Get("/formdata", middlewares.Protected(), deviceHandlers.FetchDeviceFormData)
	group.Get("/", middlewares.Protected(), deviceHandlers.FetchDevices)
	group.Post("/", middlewares.Protected(), deviceHandlers.CreateNewDevice)
}
