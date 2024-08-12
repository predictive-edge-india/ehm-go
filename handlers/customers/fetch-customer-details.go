package customerHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchCustomerDetails(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number && userRole.AccessType != models.UserRoleEnum.CustomerAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	customerIdStr := c.Params("customerId")
	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	var customer models.Customer
	err = database.Database.
		Select("id, name, logo_url, ST_AsGeoJSON(position) as position, email, phone, address_1, address_2, city, state, country, postal_code, gstin, created_at").
		Where("id = ?", customerId).Find(&customer).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching customer.")
	}

	if customer.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Customer")
	}

	payload := fiber.Map{
		"customer": customer.Json(),
	}

	return c.JSON(helpers.BuildResponse(payload))
}
