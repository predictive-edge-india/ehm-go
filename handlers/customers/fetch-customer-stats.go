package customerHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchCustomerStats(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number && userRole.AccessType != models.UserRoleEnum.CustomerAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	customerIdStr := c.Params("customerId")
	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var usersCount int64
	err = database.Database.Model(&models.UserRole{}).Preload("User").Where("customer_id = ?", customerId).Count(&usersCount).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customer user count.")
	}

	var deviceCount int64
	err = database.Database.Model(models.Device{}).
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.customer_id = ?", customerId).
		Count(&deviceCount).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customer device count.")
	}

	var assetCount int64
	err = database.Database.Model(&models.Asset{}).Where("customer_id = ?", customerId).Count(&assetCount).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customer asset count.")
	}

	payload := fiber.Map{
		"userCount":   usersCount,
		"deviceCount": deviceCount,
		"assetCount":  assetCount,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
