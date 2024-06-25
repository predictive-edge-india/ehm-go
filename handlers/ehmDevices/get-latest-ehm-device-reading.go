package ehmDeviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/helpers"
)

func GetLatestEhmDeviceReading(c *fiber.Ctx) error {
	ehmDeviceId, err := uuid.Parse(c.Params("ehmDeviceId"))
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	paramType := c.Params("paramType")

	ehmDevice := database.FindEhmDeviceById(ehmDeviceId)
	if ehmDevice.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "EHM device")
	}

	currentParameter := database.FindLatestEhmDeviceCurrentParameter(ehmDevice.Id, paramType)
	if currentParameter.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "Current parameter")
	}

	tempCoeff := database.GetTemperatureCoefficientForDevice(ehmDevice.Id)
	if tempCoeff.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "Temperature coefficient")
	}

	payload := fiber.Map{
		"currentParameter":       currentParameter.Json(),
		"temperatureCoefficient": tempCoeff.Json(),
	}
	return c.JSON(helpers.BuildResponse(payload))
}
