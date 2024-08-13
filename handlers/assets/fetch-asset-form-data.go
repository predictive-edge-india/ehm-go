package assetHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetFormData(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var assetClasses []models.AssetClass
	err = database.Database.
		Find(&assetClasses).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset classes.")
	}

	assetClassJson := make([]fiber.Map, 0)
	for _, assetClass := range assetClasses {
		assetClassJson = append(assetClassJson, fiber.Map{
			"id":   assetClass.Id,
			"name": assetClass.Name,
		})
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
		"assetClasses": assetClassJson,
		"customers":    customerJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
