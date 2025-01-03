package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

func GetPowerLimit(c *fiber.Ctx) error {
	limitTypeStr := c.Query("type")
	assetIdStr := c.Params("assetId")
	assetId, err := uuid.Parse(assetIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	powerLimit := database.FindPowerLimitByAssetIdAndLimitType(assetId, limitTypeStr)
	if powerLimit.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Power limit")
	}

	return c.JSON(helpers.BuildResponse(powerLimit.Json()))
}
