package authHandlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

func ValidateUser(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)
	if user.Id == uuid.Nil {
		return helpers.ResourceNotFoundError(c, "User")
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return helpers.NotAuthorizedError(c, "Invalid token!")
	}

	payload := fiber.Map{
		"user": user.Json(),
	}
	return c.JSON(helpers.BuildResponse(payload))
}
