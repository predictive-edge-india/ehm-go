package deviceTypeHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func DeleteDeviceType(c *fiber.Ctx) error {
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

	deviceCount := int64(0)
	err = database.Database.Model(&models.Device{}).Where("device_type_id = ?", deviceTypeId).Count(&deviceCount).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting device type.")
	}

	if deviceCount > 0 {
		return helpers.BadRequestError(c, "Device type is in use! Please delete devices first.")
	}

	var deviceTypes []models.DeviceType

	err = database.Database.Where("id = ?", deviceTypeId).Delete(&deviceTypes).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting device type.")
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"message": "Device type deleted successfully!",
	}))
}
