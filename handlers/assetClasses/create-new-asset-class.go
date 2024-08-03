package assetClassHandlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func CreateNewAssetClass(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	return validateRequestBody(c)
}

func validateRequestBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Name string `json:"name" validate:"required"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewAssetClass: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewAssetClass: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	// Check if asset class already exists
	assetClass := models.AssetClass{}
	if err := database.Database.Where("name = ?", jsonBody.Name).First(&assetClass).Error; err == nil {
		return helpers.BadRequestError(c, "Asset class already exists!")
	}

	newAsset := models.AssetClass{
		Name: jsonBody.Name,
	}

	if err := database.Database.Create(&newAsset).Error; err != nil {
		log.Error().AnErr("CreateNewAssetClass: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"assetClass": newAsset.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
