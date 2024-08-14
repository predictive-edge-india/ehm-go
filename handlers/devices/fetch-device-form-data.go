package deviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDeviceFormData(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var deviceTypes []models.DeviceType
	err = database.Database.
		Find(&deviceTypes).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching device types.")
	}

	deviceTypeJson := make([]fiber.Map, 0)
	for _, deviceType := range deviceTypes {
		deviceTypeJson = append(deviceTypeJson, fiber.Map{
			"id":   deviceType.Id,
			"name": deviceType.Name,
		})
	}

	payload := fiber.Map{
		"deviceTypes": deviceTypeJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
