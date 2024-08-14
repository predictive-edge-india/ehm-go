package assetParameterHandlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func CreateNewAssetParameter(c *fiber.Ctx) error {
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
		Name       string `json:"name" validate:"required"`
		Type       string `json:"type" validate:"required"`
		AssetClass string `json:"assetClass" validate:"required,uuid4"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewAssetParameter: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewAssetParameter: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	assetClassId, err := uuid.Parse(jsonBody.AssetClass)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid Asset Class UUID!")
	}

	// Check if device type already exists
	assetParameter := models.AssetParameter{}
	if err := database.Database.Where("name = ? AND asset_class_id = ? AND type = ?", jsonBody.Name, assetClassId, jsonBody.Type).First(&assetParameter).Error; err == nil {
		return helpers.BadRequestError(c, "Asset parameter already exists!")
	}

	newAssetParameter := models.AssetParameter{
		Name:         jsonBody.Name,
		Type:         jsonBody.Type,
		AssetClassId: &assetClassId,
	}

	if err := database.Database.Create(&newAssetParameter).Error; err != nil {
		log.Error().AnErr("CreateNewAssetParameter: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error creating asset parameter!")
	}

	payload := fiber.Map{
		"assetParameter": newAssetParameter.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
