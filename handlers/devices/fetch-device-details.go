package deviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDeviceDetails(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	deviceIdStr := c.Params("deviceId")
	deviceId, err := uuid.Parse(deviceIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var device models.Device
	err = database.Database.
		Preload("DeviceType").
		Where("id = ?", deviceId).
		Find(&device).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching device.")
	}

	if device.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Device")
	}

	assetDevice := database.FindAssetDeviceForDevice(device.Id)

	payload := device.Json()
	if !assetDevice.IsIdNull() {
		payload["asset"] = assetDevice.Asset.ShortJson()
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"device": payload,
	}))
}
