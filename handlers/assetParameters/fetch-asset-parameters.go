package assetParameterHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchAssetParameters(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	searchQuery := strings.Trim(c.Query("q"), " ")
	typeQuery := strings.Trim(c.Query("type"), " ")
	assetClassIdStr := strings.Trim(c.Query("asset_class"), " ")

	var assetParameters []models.AssetParameter
	query := database.Database

	if len(searchQuery) > 0 {
		query = query.Where("name ILIKE ?", "%"+searchQuery+"%")
	}

	if len(typeQuery) > 0 {
		query = query.Where("type = ?", typeQuery)
	}

	assetClassId, err := uuid.Parse(assetClassIdStr)
	if err == nil {
		query = query.Where("asset_class_id = ?", assetClassId)
	}

	err = query.
		Order("param_order asc").
		Find(&assetParameters).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching asset parameters.")
	}

	var payload = make([]fiber.Map, 0)
	for _, param := range assetParameters {
		payload = append(payload, param.Json())
	}

	return c.JSON(helpers.BuildResponse(fiber.Map{
		"list": payload,
	}))
}
