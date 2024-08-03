package deviceTypeHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchDeviceTypes(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)
	page, perPage := helpers.GetPagination(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	searchQuery := strings.Trim(c.Query("q"), " ")

	var deviceTypes []models.DeviceType

	query := database.Database

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}
	err = query.
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&deviceTypes).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching device types.")
	}

	var total int64
	if err = query.Model(&models.DeviceType{}).Count(&total).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error counting device types.")
	}

	var payload = make([]fiber.Map, 0)
	for _, deviceType := range deviceTypes {
		payload = append(payload, deviceType.ShortJson())
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
