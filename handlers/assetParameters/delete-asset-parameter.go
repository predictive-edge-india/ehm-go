package assetParameterHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func DeleteAssetParameter(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	assetParameterIdStr := c.Params("assetParameterId")
	assetParameterId, err := uuid.Parse(assetParameterIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var assetParameter models.AssetParameter

	err = database.Database.Unscoped().Where("id = ?", assetParameterId).Delete(&assetParameter).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting asset parameter.")
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"message": "Asset parameter deleted successfully!",
	}))
}
