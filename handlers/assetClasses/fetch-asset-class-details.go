package assetClassHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetClassDetails(c *fiber.Ctx) error {
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

	var assetClass models.AssetClass
	err = database.Database.
		Where("id = ?", assetClassId).Find(&assetClass).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset class.")
	}

	if assetClass.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset class")
	}

	payload := fiber.Map{
		"assetClass": assetClass.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
