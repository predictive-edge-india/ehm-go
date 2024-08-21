package assetMonitoringHandlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetPowerParameterData(c *fiber.Ctx) error {
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

	paramKeyStr := c.Params("paramKey")
	if len(paramKeyStr) == 0 {
		return helpers.BadRequestError(c, "Invalid parameter key!")
	}

	asset := database.FindAssetByIdWithAssetClass(assetId)
	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	instantData := database.GetPowerParameterForAsset(asset.Id, paramKeyStr)

	// date range
	startDateStr := c.Query("start", "")
	endDateStr := c.Query("end", "")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		startDate = time.Now().AddDate(0, 0, -1)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		endDate = time.Now()
	}

	paramDataList := database.GetPowerParameterForAssetWithRange(asset.Id, paramKeyStr, startDate, endDate)

	paramDataListJson := make([]fiber.Map, 0)
	for _, paramData := range paramDataList {
		paramDataListJson = append(paramDataListJson, fiber.Map{
			"timestamp": paramData.Time,
			"value":     paramData.Value,
		})
	}

	payload := fiber.Map{
		"instantData": instantData.Value,
		"paramData":   paramDataListJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
