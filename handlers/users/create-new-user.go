package userHandlers

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewUser(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	log.Info().Any("user", user).Send()
	currentCustomer, customerRequest, err := database.FindCurrentUserCustomer(c, user)
	if err != nil && customerRequest {
		return err
	}

	userRole, err := database.FindUserRoleForCustomerUser(c, user, currentCustomer)
	if err != nil {
		return err
	}

	if userRole.AccessType == models.UserRoleEnum.SuperAdministrator.Number {
		return validateBody(c)
	}

	return helpers.NotAuthorizedError(c, "You're not authorized!")
}

func validateBody(c *fiber.Ctx) error {
	jsonBody := struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone"`
		Password string `json:"password" validate:"required,min=6"`
		Role     int16  `json:"role" validate:"required"`
		Customer string `json:"customer" validate:"omitempty,uuid4"`
	}{}

	// Validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("CreateNewUser: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		log.Error().AnErr("CreateNewUser: Validator", err).Send()
		return helpers.BadRequestError(c, "Please check your request!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().AnErr("CreateNewUser: GenerateFromPassword", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	userRole := models.UserRole{
		AccessType: jsonBody.Role,
	}

	if jsonBody.Role != models.UserRoleEnum.SuperAdministrator.Number {
		customerId, err := uuid.Parse(jsonBody.Customer)
		if err != nil {
			log.Error().AnErr("CreateNewUser: UUID parsing", err).Send()
			return helpers.BadRequestError(c, "Invalid Customer UUID!")
		}

		customer := database.FindCustomerById(customerId)
		if customer.IsIdNull() {
			return helpers.ResourceNotFoundError(c, "Customer")
		}

		userRole.CustomerId = &customerId
	}

	newCustomer := models.User{
		Name:         jsonBody.Name,
		Email:        jsonBody.Email,
		Phone:        sql.NullString{String: jsonBody.Phone, Valid: jsonBody.Phone != ""},
		PasswordHash: string(hashedPassword),
	}

	transaction := database.Database.Begin()
	if err := transaction.Create(&newCustomer).Error; err != nil {
		log.Error().AnErr("CreateNewUser: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	userRole.UserId = newCustomer.Id

	if err := transaction.Create(&userRole).Error; err != nil {
		log.Error().AnErr("CreateNewUser: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	err = transaction.Commit().Error
	if err != nil {
		log.Error().AnErr("CreateNewUser: Database", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	payload := fiber.Map{
		"user": newCustomer.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
