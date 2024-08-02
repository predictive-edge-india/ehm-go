package ehmDeviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iisc/demo-go/helpers"
)

func FetchAllCustomers(c *fiber.Ctx) error {

	return c.JSON(helpers.BuildResponse("payload"))
}
