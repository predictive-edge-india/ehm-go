package deviceTypeHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDeviceTypeDetails(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	deviceTypeIdStr := c.Params("deviceTypeId")
	deviceTypeId, err := uuid.Parse(deviceTypeIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var deviceType models.DeviceType
	err = database.Database.
		Where("id = ?", deviceTypeId).Find(&deviceType).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching device type.")
	}

	if deviceType.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Device type")
	}

	payload := fiber.Map{
		"deviceType": deviceType.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
