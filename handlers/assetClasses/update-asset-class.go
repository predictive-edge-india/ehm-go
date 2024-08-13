package assetClassHandlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func UpdateAssetClass(c *fiber.Ctx) error {
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

	jsonBody := struct {
		Name string `json:"name" validate:"required"`
	}{}

	if err := c.BodyParser(&jsonBody); err != nil {
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		return helpers.BadRequestError(c, "Please check your request!")
	}

	assetClass.Name = jsonBody.Name
	err = database.Database.Save(&assetClass).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"assetClass": assetClass.Json(),
	}
	return c.JSON(helpers.BuildResponse(payload))
}
