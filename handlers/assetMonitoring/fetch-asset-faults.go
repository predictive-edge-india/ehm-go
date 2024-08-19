package assetMonitoringHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetFaults(c *fiber.Ctx) error {
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

	asset := database.FindAssetByIdWithAssetClass(assetId)
	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	var assetParameters []models.AssetParameter
	database.Database.Where("asset_class_id = ?", asset.AssetClass.Id).Find(&assetParameters)

	assetStatusFlag := database.GetStatusFlagForAsset(asset.Id)
	if assetStatusFlag.Id == uuid.Nil {
		return helpers.BadRequestError(c, "Asset faults not recorded!")
	}

	var assetParametersJson = make([]fiber.Map, 0)
	for index, assetParameter := range assetParameters {
		assetParametersJson = append(assetParametersJson, fiber.Map{
			"id":     assetParameter.Id,
			"name":   assetParameter.Name,
			"status": assetStatusFlag.Statuses[index] == 1,
		})
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"assetFaults": assetParametersJson,
	}))
}
