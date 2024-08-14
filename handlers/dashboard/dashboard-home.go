package dashboardHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDashboardHome(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var result struct {
		Users     int
		Devices   int
		Assets    int
		Customers int
	}

	database.Database.Raw(`
		SELECT
			COALESCE(
				(
					SELECT COUNT(*)
					FROM users
				),
				0
			) AS users,
			COALESCE(
					(
						SELECT COUNT(*)
						FROM devices
					),
					0
				) AS devices,
			COALESCE(
					(
						SELECT COUNT(*)
						FROM assets
					),
					0
				) AS assets,
			COALESCE(
					(
						SELECT COUNT(*)
						FROM customers
					),
					0
				) AS customers
	`).Scan(&result)

	payload := fiber.Map{
		"users":     result.Users,
		"devices":   result.Devices,
		"assets":    result.Assets,
		"customers": result.Customers,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
