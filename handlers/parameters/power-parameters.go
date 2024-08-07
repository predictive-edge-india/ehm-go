package parameterHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

func GetPowerParameter(c *fiber.Ctx) error {
	ehmDeviceId, err := uuid.Parse(c.Params("ehmDeviceId"))
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	paramType := c.Params("paramType")

	ehmDevice := database.FindEhmDeviceById(ehmDeviceId)
	if ehmDevice.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "EHM device")
	}

	powerParam := database.GetPowerParameterForDevice(ehmDevice.Id, paramType)
	if powerParam.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "Fault panel")
	}

	payload := fiber.Map{
		"powerParameter": powerParam.Json(),
	}
	return c.JSON(helpers.BuildResponse(payload))
}
