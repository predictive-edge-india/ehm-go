package assetMonitoringHandlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

type PowerDataParamValuesWithTime struct {
	Values []string `gorm:"column:values"`
	Time   string   `gorm:"column:created_at"`
}

func FetchAssetLoadImbalanceData(c *fiber.Ctx) error {
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

	loadL1DataList := database.GetPowerParameterForAssetWithRange(asset.Id, "load_l1_current", startDate, endDate)
	loadL2DataList := database.GetPowerParameterForAssetWithRange(asset.Id, "load_l2_current", startDate, endDate)
	loadL3DataList := database.GetPowerParameterForAssetWithRange(asset.Id, "load_l3_current", startDate, endDate)

	consolidatedData := make([]PowerDataParamValuesWithTime, 0)
	for index, loadL1Data := range loadL1DataList {
		if (loadL1Data.Time == loadL2DataList[index].Time) && (loadL1Data.Time == loadL3DataList[index].Time) {
			consolidatedData = append(consolidatedData, PowerDataParamValuesWithTime{
				Values: []string{
					loadL1Data.Value,
					loadL2DataList[index].Value,
					loadL3DataList[index].Value,
				},
				Time: loadL1Data.Time,
			})
		}
	}

	paramDataListJson := make([]fiber.Map, 0)
	for _, paramData := range consolidatedData {
		paramDataListJson = append(paramDataListJson, fiber.Map{
			"timestamp": paramData.Time,
			"values":    paramData.Values,
		})
	}

	payload := fiber.Map{
		"paramData": paramDataListJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
