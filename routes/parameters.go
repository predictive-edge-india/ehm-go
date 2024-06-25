package routes

import (
	"github.com/gofiber/fiber/v2"
	parameterHandlers "github.com/iisc/demo-go/handlers/parameters"
)

func ParameterRoutes(router fiber.Router) {
	devices := router.Group("/devices/:ehmDeviceId")
	parameters := devices.Group("/parameters")
	parameters.Get("/fault-panel", parameterHandlers.GetFaultPanel)
	parameters.Get("/power-parameter", parameterHandlers.GetPowerParameter)
	parameters.Get("/engine-data", parameterHandlers.GetEngineData)
}
