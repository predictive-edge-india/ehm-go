package assetMonitoringHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetLastLocation(c *fiber.Ctx) error {
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

	lastLocation := database.GetLastLocationForAsset(asset.Id)
	if lastLocation.Id == uuid.Nil {
		return helpers.BadRequestError(c, "Last location not recorded!")
	}

	payload := fiber.Map{
		"lastLocation": lastLocation.ShortJson(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
