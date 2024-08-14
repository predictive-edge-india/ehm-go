package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func DeleteUser(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	userIdStr := c.Params("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	err = database.Database.Where("id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error deleting user.")
	}

	return c.JSON(helpers.BuildResponse("Success"))
}
