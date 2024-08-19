package dashboardHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchCustomerLocations(c *fiber.Ctx) error {
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
		Select("id, name, ST_AsGeoJSON(position) as position").
		Find(&customers).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customers.")
	}

	var customerJson []fiber.Map
	for _, customer := range customers {
		customerJson = append(customerJson, fiber.Map{
			"id":   customer.Id,
			"name": customer.Name,
			"position": fiber.Map{
				"type":        "Point",
				"coordinates": customer.Position.Coordinates,
			},
		})
	}

	payload := fiber.Map{
		"customers": customerJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
