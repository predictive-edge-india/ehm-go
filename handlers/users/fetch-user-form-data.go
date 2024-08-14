package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchUserFormData(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var customers []models.Customer
	err = database.Database.
		Select("id, name").
		Find(&customers).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customers.")
	}

	customerJson := make([]fiber.Map, 0)
	for _, customer := range customers {
		customerJson = append(customerJson, fiber.Map{
			"id":   customer.Id,
			"name": customer.Name,
		})
	}

	payload := fiber.Map{
		"customers": customerJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
