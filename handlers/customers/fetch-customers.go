package customerHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func FetchAllCustomers(c *fiber.Ctx) error {
	// user := database.FindUserAuth(c)
	page, perPage := helpers.GetPagination(c)

	searchQuery := strings.Trim(c.Query("q"), " ")

	query := database.Database

	var customers []models.Customer

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}

	err := query.
		Select("id, name, logo_url, ST_AsGeoJSON(position) as position, created_at").
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&customers).Error
	if err != nil {
		log.Error().AnErr("Fetch customers", err).Send()
		return helpers.BadRequestError(c, "There was an error fetching customers.")
	}

	var total int64
	if err = query.Model(&models.Customer{}).Count(&total).Error; err != nil {
		log.Error().AnErr("Count customers", err).Send()
		return helpers.BadRequestError(c, "There was an error counting customers.")
	}

	var payload = make([]fiber.Map, 0)
	for _, customer := range customers {
		payload = append(payload, customer.ShortJson())
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
