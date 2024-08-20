package assetMonitoringHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetDGStatus(c *fiber.Ctx) error {
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

	dgStatus := database.GetDGStatusForAsset(asset.Id)
	if dgStatus.Id == uuid.Nil {
		return helpers.BadRequestError(c, "DG status not recorded!")
	}

	payload := fiber.Map{
		"dgStatus": dgStatus.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
