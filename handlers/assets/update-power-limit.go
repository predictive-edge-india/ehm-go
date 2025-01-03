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

func UpdatePowerLimit(c *fiber.Ctx) error {
	assetIdStr := c.Params("assetId")
	assetId, err := uuid.Parse(assetIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	jsonBody := struct {
		UpperLimit int16  `json:"upperLimit" validate:"required"`
		LowerLimit int16  `json:"lowerLimit" validate:"required"`
		LimitType  string `json:"type" validate:"required"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("UpdatePowerLimit: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("UpdatePowerLimit: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	powerLimit := database.FindPowerLimitByAssetIdAndLimitType(assetId, jsonBody.LimitType)
	if powerLimit.IsIdNull() {
		err := createPowerLimit(assetId, jsonBody.LimitType, jsonBody.UpperLimit, jsonBody.LowerLimit)
		if err != nil {
			return helpers.BadRequestError(c, err)
		}
	}

	powerLimit.UpperLimit = jsonBody.UpperLimit
	powerLimit.LowerLimit = jsonBody.LowerLimit
	err = database.Database.Save(&powerLimit).Error
	if err != nil {
		return helpers.BadRequestError(c, err)
	}

	return c.JSON(helpers.BuildResponse("Success"))
}

func createPowerLimit(assetId uuid.UUID, limitType string, upper, lower int16) error {
	powerLimit := models.PowerLimit{
		AssetId:    &assetId,
		Type:       limitType,
		UpperLimit: upper,
		LowerLimit: lower,
	}

	err := database.Database.Create(&powerLimit).Error
	if err != nil {
		return err
	}

	return nil
}
