package deviceHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDevices(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)
	page, perPage := helpers.GetPagination(c)

	_, customerRequest, err := database.FindCurrentUserCustomer(c, user)
	if err != nil && customerRequest {
		return err
	}

	searchQuery := strings.Trim(c.Query("q"), " ")
	deviceTypeId := strings.Trim(c.Query("device_type"), " ")
	// customerId := strings.Trim(c.Query("customer"), " ")

	var devices []models.Device

	query := database.Database

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}

	if len(deviceTypeId) > 0 {
		query = query.Where("device_type_id = ?", deviceTypeId)
	}

	err = query.
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&devices).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching devices.")
	}

	var total int64
	if err = query.Model(&models.Device{}).Count(&total).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error counting devices.")
	}

	var payload = make([]fiber.Map, 0)
	for _, device := range devices {
		payload = append(payload, device.ShortJson())
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"list": payload,
		"_meta": fiber.Map{
			"perPage": perPage,
			"page":    page,
			"total":   total,
		},
	}))
}
