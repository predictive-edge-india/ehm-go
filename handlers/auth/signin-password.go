package authHandlers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/middlewares"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func SigninWithPassword(c *fiber.Ctx) error {
	jsonBody := struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

	//validation
	if err := c.BodyParser(&jsonBody); err != nil {
		log.Error().AnErr("SigninWithPassword: Bodyparser", err).Send()
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	err := validate.Struct(jsonBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.StructNamespace() == "Password" && err.Tag() == "min" {
				return helpers.BadRequestError(c, "Password should be minimum 6 characters long")
			}
		}
	}

	user, err := database.FindUserByEmail(jsonBody.Email)
	if err != nil {
		log.Error().AnErr("SigninWithPassword: FindUserByEmail", err).Send()
		return helpers.BadRequestError(c, err.Error())
	}

	if user.IsIdNull() {
		return helpers.BadRequestError(c, "Email/password is incorrect!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(jsonBody.Password))
	if err != nil {
		log.Error().AnErr("SigninWithPassword: CompareHashAndPassword", err).Send()
		return helpers.BadRequestError(c, "Email/password is incorrect!")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(middlewares.SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	payload := fiber.Map{
		"user":  user.Json(),
		"token": t,
	}
	return c.JSON(helpers.BuildResponse(payload))
}
