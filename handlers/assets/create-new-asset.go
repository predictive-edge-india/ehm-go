package assetHandlers

import (
	"database/sql"

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
		Make       string `json:"make" validate:"required"`
		Model      string `json:"model" validate:"required"`
		Name       string `json:"name"`
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

	customer := database.FindCustomerById(customerId)
	if customer.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Customer")
	}

	assetClassId, err := uuid.Parse(jsonBody.AssetClass)
	if err != nil {
		log.Error().AnErr("CreateNewAsset: UUID parsing", err).Send()
		return helpers.BadRequestError(c, "Invalid Asset Class UUID!")
	}

	assetClass := database.FindAssetClassById(assetClassId)
	if assetClass.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset Class")
	}

	newAsset := models.Asset{
		Make:         jsonBody.Make,
		ModelName:    jsonBody.Model,
		Name:         sql.NullString{String: jsonBody.Name, Valid: len(jsonBody.Name) > 0},
		CustomerId:   &customerId,
		AssetClassId: &assetClassId,
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
