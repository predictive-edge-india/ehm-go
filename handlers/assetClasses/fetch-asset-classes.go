package assetClassHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetClasses(c *fiber.Ctx) error {
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

	var assets []models.AssetClass

	query := database.Database

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}
	err = query.
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&assets).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset classes.")
	}

	var total int64
	if err = query.Model(&models.AssetClass{}).Count(&total).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error counting asset classes.")
	}

	var payload = make([]fiber.Map, 0)
	for _, asset := range assets {
		payload = append(payload, asset.ShortJson())
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
