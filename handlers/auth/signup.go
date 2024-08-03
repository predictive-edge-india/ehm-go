package authHandlers

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(c *fiber.Ctx) error {
	jsonBody := struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
	}{}
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("SignupUser: bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	err := validate.Struct(jsonBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.StructNamespace() == "Email" && err.Tag() == "email" {
				return helpers.BadRequestError(c, "Please check the email")
			}
			if err.StructNamespace() == "Password" && err.Tag() == "min" {
				return helpers.BadRequestError(c, "Password should be minimum 6 characters long")
			}
		}
		return helpers.BadRequestError(c, "Please check your request!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonBody.Password), 8)
	if err != nil {
		log.Error().AnErr("SignupUser: GenerateFromPassword", err).Send()
		return helpers.BadRequestError(c, "There was an error!")
	}

	user := new(models.User)

	user.Name = jsonBody.Name
	user.Email = jsonBody.Email
	user.Phone = sql.NullString{
		Valid:  true,
		String: jsonBody.Phone,
	}
	user.PasswordHash = string(hashedPassword)

	var count int64
	database.Database.Model(&models.User{}).Where("email = ?", jsonBody.Email).Or("phone = ?", jsonBody.Phone).Count(&count)

	if count > 0 {
		return helpers.BadRequestError(c, "Email or Phone already exists!")
	}

	if err = database.Database.Create(user).Error; err != nil {
		log.Error().AnErr("SignupUser: CreateUserDB", err).Send()
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"user": user.Json(),
	}))
}
