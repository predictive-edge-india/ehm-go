package assetClassHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func DeleteAssetClass(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	assetClassIdStr := c.Params("assetClassId")
	assetClassId, err := uuid.Parse(assetClassIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	assetCount := int64(0)
	err = database.Database.Model(&models.Asset{}).Where("asset_class_id = ?", assetClassId).Count(&assetCount).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting asset class.")
	}

	if assetCount > 0 {
		return helpers.BadRequestError(c, "Asset class is in use! Please delete assets first.")
	}

	var assets []models.AssetClass

	err = database.Database.Where("id = ?", assetClassId).Delete(&assets).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting asset class.")
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"message": "Asset class deleted successfully!",
	}))
}
