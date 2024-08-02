package ehmDeviceHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

func FetchAllCustomers(c *fiber.Ctx) error {

	return c.JSON(helpers.BuildResponse("payload"))
}
