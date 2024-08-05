package profileHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchProfileCustomer(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	var userRoles []models.UserRole
	database.Database.
		Preload("Customer").
		Where("user_id = ?", user.Id).
		Find(&userRoles)

	payload := make([]fiber.Map, 0)
	for _, userRole := range userRoles {
		if userRole.AccessType == models.UserRoleEnum.SuperAdministrator.Number {
			payload = append(make([]fiber.Map, 0), fiber.Map{
				"accessType": userRole.AccessType,
			})
			return c.JSON(helpers.BuildResponse(fiber.Map{
				"customers": payload,
			}))
		} else {
			orgJson := userRole.Customer.Json()
			orgJson["accessType"] = userRole.AccessType
			payload = append(payload, orgJson)
		}

	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"customers": payload,
	}))
}
