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

func UpdateAssetDetails(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	currentCustomer, requestCustomer, err := database.FindCurrentUserCustomer(c, user)
	if err != nil && requestCustomer {
		return err
	}

	userRole, err := database.FindUserRoleForCustomerUser(c, user, currentCustomer)
	if err != nil {
		return err
	}

	if userRole.AccessType == models.UserRoleEnum.SuperAdministrator.Number {
		return validateAssetUpdateBody(c)
	}

	return helpers.NotAuthorizedError(c, "You're not authorized!")
}

func validateAssetUpdateBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Make  string `json:"make"`
		Model string `json:"model"`
		Name  string `json:"name"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("UpdateAssetDetails: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("UpdateAssetDetails: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	assetIdStr := c.Params("assetId")
	assetId, err := uuid.Parse(assetIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid asset UUID!")
	}

	asset := database.FindAssetById(assetId)
	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	if jsonBody.Make != "" {
		asset.Make = jsonBody.Make
	}

	if jsonBody.Model != "" {
		asset.ModelName = jsonBody.Model
	}

	if jsonBody.Name != "" {
		asset.Name = jsonBody.Name
	}

	if err := database.Database.Save(&asset).Error; err != nil {
		log.Error().AnErr("UpdateAssetDetails: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"asset": asset.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
