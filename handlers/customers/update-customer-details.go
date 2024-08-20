package customerHandlers

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

func UpdateCustomerDetails(c *fiber.Ctx) error {
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
		return validateCustomerUpdateBody(c)
	}

	return helpers.NotAuthorizedError(c, "You're not authorized!")
}

func validateCustomerUpdateBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Name       string         `json:"name"`
		Email      string         `json:"email" validate:"email"`
		Phone      string         `json:"phone"`
		LogoUrl    string         `json:"logoUrl"`
		Position   models.GeoJson `json:"position"`
		Address1   string         `json:"address1"`
		Address2   string         `json:"address2"`
		City       string         `json:"city"`
		State      string         `json:"state"`
		Country    string         `json:"country"`
		PostalCode int32          `json:"postalCode"`
		Gstin      string         `json:"gstin"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("UpdateCustomer: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("UpdateCustomer: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	customerIdStr := c.Params("customerId")
	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	customer := database.FindCustomerById(customerId)
	if customer.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Customer")
	}

	if jsonBody.Name != "" {
		customer.Name = jsonBody.Name
	}

	if jsonBody.Email != "" {
		customer.Email = jsonBody.Email
	}

	if jsonBody.Phone != "" {
		customer.Phone = jsonBody.Phone
	}

	if jsonBody.LogoUrl != "" {
		customer.LogoUrl = sql.NullString{String: jsonBody.LogoUrl, Valid: jsonBody.LogoUrl != ""}
	}

	// Name:       jsonBody.Name,
	// LogoUrl:    sql.NullString{String: jsonBody.LogoUrl, Valid: jsonBody.LogoUrl != ""},
	// Position:   jsonBody.Position,
	// Address1:   jsonBody.Address1,
	// Address2:   sql.NullString{String: jsonBody.Address2, Valid: jsonBody.Address2 != ""},
	// City:       jsonBody.City,
	// State:      jsonBody.State,
	// Country:    jsonBody.Country,
	// PostalCode: jsonBody.PostalCode,

	if err := database.Database.Save(&customer).Error; err != nil {
		log.Error().AnErr("UpdateCustomer: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"customer": customer.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
