package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"gorm.io/gorm"
)

func FetchAssetDetails(c *fiber.Ctx) error {
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

	var asset models.Asset
	err = database.Database.
		Preload("AssetClass").
		Preload("Customer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Where("id = ?", assetId).
		Find(&asset).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset.")
	}

	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	payload := fiber.Map{
		"asset": asset.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
