package parameterHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/helpers"
)

func GetFaultPanel(c *fiber.Ctx) error {
	ehmDeviceId, err := uuid.Parse(c.Params("ehmDeviceId"))
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	paramType := c.Params("paramType")

	ehmDevice := database.FindEhmDeviceById(ehmDeviceId)
	if ehmDevice.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "EHM device")
	}

	faultPanel := database.GetFaultPanelForDevice(ehmDevice.Id, paramType)
	if faultPanel.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "Fault panel")
	}

	payload := fiber.Map{
		"faultPanel": faultPanel.Json(),
	}
	return c.JSON(helpers.BuildResponse(payload))
}
