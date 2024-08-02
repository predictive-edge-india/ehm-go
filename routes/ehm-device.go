package routes

import (
	"github.com/gofiber/fiber/v2"
	ehmDeviceHandlers "github.com/predictive-edge-india/ehm-go/handlers/ehmDevices"
)

func EhmDeviceRoutes(app fiber.Router) {
	devices := app.Group("/devices")
	devices.Get("/:ehmDeviceId/:paramType", ehmDeviceHandlers.GetLatestEhmDeviceReading)
	devices.Get("/", ehmDeviceHandlers.FetchAllEhmDevices)
}
