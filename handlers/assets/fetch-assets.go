package assetHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssets(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)
	page, perPage := helpers.GetPagination(c)

	currentCustomer, customerRequest, err := database.FindCurrentUserCustomer(c, user)
	if err != nil && customerRequest {
		return err
	}

	searchQuery := strings.Trim(c.Query("q"), " ")
	assetClassId := strings.Trim(c.Query("asset_class"), " ")
	customerId := strings.Trim(c.Query("customer"), " ")

	var assets []models.Asset

	query := database.Database

	if !currentCustomer.IsIdNull() {
		query = query.Where("customer_id = ?", currentCustomer.Id)
	}

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}

	if len(assetClassId) > 0 {
		query = query.Where("asset_class_id = ?", assetClassId)
	}

	if len(customerId) > 0 {
		query = query.Where("customer_id = ?", customerId)
	}

	err = query.
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&assets).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching assets.")
	}

	var total int64
	if err = query.Model(&models.Asset{}).Count(&total).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error counting assets.")
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
