package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func UnassignAssetDevice(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	assetIdStr := c.Params("assetId")
	assetId, err := uuid.Parse(assetIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid asset UUID!")
	}

	asset := database.FindAssetById(assetId)
	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	deviceIdStr := c.Params("deviceId")
	deviceId, err := uuid.Parse(deviceIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid device UUID!")
	}

	assetDevice := database.FindAssetDevice(assetId, deviceId)
	if assetDevice.IsIdNull() {
		return helpers.BadRequestError(c, "Device was not assigned to asset!")
	}

	if err := database.Database.Unscoped().Delete(&assetDevice).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error unassigning device from asset!")
	}

	return c.JSON(helpers.BuildResponse("Success"))
}
