package customerHandlers

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func CreateNewCustomer(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	currentCustomer, _, err := database.FindCurrentUserCustomer(c, user)
	if err != nil {
		return err
	}

	userRole, err := database.FindUserRoleForCustomerUser(c, user, currentCustomer)
	if err != nil {
		return err
	}

	if userRole.AccessType == models.UserRoleEnum.SuperAdministrator.Number {
		return validateCustomerBody(c)
	}

	return helpers.NotAuthorizedError(c, "You're not authorized!")
}

func validateCustomerBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Name       string         `json:"name" validate:"required"`
		LogoUrl    string         `json:"logoUrl"`
		Position   models.GeoJson `json:"position" validate:"required"`
		Address1   string         `json:"address1" validate:"required"`
		Address2   string         `json:"address2"`
		City       string         `json:"city" validate:"required"`
		State      string         `json:"state" validate:"required"`
		Country    string         `json:"country" validate:"required"`
		PostalCode string         `json:"postalCode" validate:"required"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewCustomer: bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewCustomer: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	newCustomer := models.Customer{
		Name:       jsonBody.Name,
		LogoUrl:    sql.NullString{String: jsonBody.LogoUrl, Valid: jsonBody.LogoUrl != ""},
		Position:   jsonBody.Position,
		Address1:   jsonBody.Address1,
		Address2:   sql.NullString{String: jsonBody.Address2, Valid: jsonBody.Address2 != ""},
		City:       jsonBody.City,
		State:      jsonBody.State,
		Country:    jsonBody.Country,
		PostalCode: jsonBody.PostalCode,
	}

	if err := database.Database.Create(&newCustomer).Error; err != nil {
		log.Error().AnErr("CreateNewCustomer: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"customer": newCustomer.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
