package deviceHandlers

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

func CreateNewDevice(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	return validateDeviceBody(c)
}

func validateDeviceBody(c *fiber.Ctx) error {
	jsonBody := struct {
		SerialNo string `json:"serialNo" validate:"required"`

		Imei1  string `json:"imei1" validate:"required"`
		Phone1 string `json:"phone1" validate:"required"`

		Imei2  string `json:"imei2"`
		Phone2 string `json:"phone2"`

		Note string `json:"note"`

		DeviceType string `json:"deviceType" validate:"required,uuid4,required"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewDevice: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewDevice: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	deviceTypeId, err := uuid.Parse(jsonBody.DeviceType)
	if err != nil {
		log.Error().AnErr("CreateNewDevice: UUID parsing", err).Send()
		return helpers.BadRequestError(c, "Invalid Device Type UUID!")
	}

	newDevice := models.Device{
		SerialNo:     jsonBody.SerialNo,
		DeviceTypeId: deviceTypeId,

		Imei1:  jsonBody.Imei1,
		Phone1: jsonBody.Phone1,

		Imei2: sql.NullString{
			String: jsonBody.Imei2,
			Valid:  jsonBody.Imei2 != "",
		},
		Phone2: sql.NullString{
			String: jsonBody.Phone2,
			Valid:  jsonBody.Phone2 != "",
		},

		Note: sql.NullString{
			String: jsonBody.Note,
			Valid:  jsonBody.Note != "",
		},
	}

	transaction := database.Database.Begin()

	if err := transaction.Create(&newDevice).Error; err != nil {
		log.Error().AnErr("CreateNewDevice: create device", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	if err := transaction.Commit().Error; err != nil {
		log.Error().AnErr("CreateNewDevice: commit", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"device": newDevice.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
