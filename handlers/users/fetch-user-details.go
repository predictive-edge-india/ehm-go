package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchUserDetails(c *fiber.Ctx) error {
	loggedUser := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, loggedUser)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number && userRole.AccessType != models.UserRoleEnum.CustomerAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	userIdStr := c.Params("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid User UUID!")
	}

	user := database.FindUserById(userId)
	if user.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "User")
	}

	payload := fiber.Map{
		"user": user.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
