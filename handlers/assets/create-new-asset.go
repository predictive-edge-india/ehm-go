package assetHandlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func CreateNewAsset(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	return validateAssetBody(c)
}

func validateAssetBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Name       string `json:"name" validate:"required"`
		Customer   string `json:"customer" validate:"required,uuid4"`
		AssetClass string `json:"assetClass" validate:"required,uuid4"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewAsset: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewAsset: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	customerId, err := uuid.Parse(jsonBody.Customer)
	if err != nil {
		log.Error().AnErr("CreateNewAsset: UUID parsing", err).Send()
		return helpers.BadRequestError(c, "Invalid Customer UUID!")
	}

	newAsset := models.Asset{
		Name:       jsonBody.Name,
		CustomerId: &customerId,
	}

	if err := database.Database.Create(&newAsset).Error; err != nil {
		log.Error().AnErr("CreateNewAsset: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"asset": newAsset.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
