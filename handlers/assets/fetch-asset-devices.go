package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetDevices(c *fiber.Ctx) error {
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
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var assetDevices []models.AssetDevice
	err = database.Database.
		Preload("Device").
		Where("asset_id = ?", assetId).
		Find(&assetDevices).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset devices.")
	}

	deviceJson := []fiber.Map{}
	for _, assetDevice := range assetDevices {
		deviceJson = append(deviceJson, fiber.Map{
			"id":       assetDevice.Device.Id,
			"name":     assetDevice.Device.Name,
			"serialNo": assetDevice.Device.SerialNo,
		})
	}

	payload := fiber.Map{
		"devices": deviceJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
