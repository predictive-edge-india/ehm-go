package ehmDeviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

func FetchAllEhmDevices(c *fiber.Ctx) error {

	devices := database.FindAllEhmDevices()

	var devicesPayload = make([]fiber.Map, 0)
	for _, device := range devices {
		obj := device.Json()
		devicesPayload = append(devicesPayload, obj)
	}

	payload := fiber.Map{
		"devices": devicesPayload,
	}
	return c.JSON(helpers.BuildResponse(payload))
}
