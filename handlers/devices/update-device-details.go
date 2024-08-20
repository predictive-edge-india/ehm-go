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

func UpdateDeviceDetails(c *fiber.Ctx) error {
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
		return validateDeviceUpdateBody(c)
	}

	return helpers.NotAuthorizedError(c, "You're not authorized!")
}

func validateDeviceUpdateBody(c *fiber.Ctx) error {
	jsonBody := struct {
		SerialNo string `json:"serialNo"`

		Imei1  string `json:"imei1"`
		Phone1 string `json:"phone1"`

		Imei2  string `json:"imei2"`
		Phone2 string `json:"phone2"`

		Note string `json:"note"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("UpdateDeviceDetails: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("UpdateDeviceDetails: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	deviceIdStr := c.Params("deviceId")
	deviceId, err := uuid.Parse(deviceIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid device UUID!")
	}

	device := database.FindDeviceById(deviceId)
	if device.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Device")
	}

	if jsonBody.SerialNo != "" {
		device.SerialNo = jsonBody.SerialNo
	}

	if jsonBody.Imei1 != "" {
		device.Imei1 = jsonBody.Imei1
	}

	if jsonBody.Phone1 != "" {
		device.Phone1 = jsonBody.Phone1
	}

	if jsonBody.Imei2 != "" {
		device.Imei2 = sql.NullString{
			String: jsonBody.Imei2,
			Valid:  true,
		}
	}

	if jsonBody.Phone2 != "" {
		device.Phone2 = sql.NullString{
			String: jsonBody.Phone2,
			Valid:  true,
		}
	}

	if jsonBody.Note != "" {
		device.Note = sql.NullString{
			String: jsonBody.Note,
			Valid:  true,
		}
	}

	if err := database.Database.Save(&device).Error; err != nil {
		log.Error().AnErr("UpdateDeviceDetails: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"device": device.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
