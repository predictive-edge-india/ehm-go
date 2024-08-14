package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssignAssetDeviceFormData(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var unassignedDevices []models.Device
	err = database.Database.
		Joins("LEFT JOIN asset_devices ON devices.id = asset_devices.device_id").
		Where("asset_devices.device_id IS NULL").
		Find(&unassignedDevices).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching unassigned devices.")
	}

	unassignedDeviceJson := make([]fiber.Map, 0)
	for _, unassignedDevice := range unassignedDevices {
		unassignedDeviceJson = append(unassignedDeviceJson, fiber.Map{
			"id":       unassignedDevice.Id,
			"name":     unassignedDevice.Name,
			"serialNo": unassignedDevice.SerialNo,
		})
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"devices": unassignedDeviceJson,
	}))
}
