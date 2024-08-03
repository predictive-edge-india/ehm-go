package userHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FetchUsers(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)
	page, perPage := helpers.GetPagination(c)

	currentCustomer, err := database.FindCurrentUserCustomer(c, user)
	if err != nil {
		return err
	}

	searchQuery := strings.Trim(c.Query("q"), " ")

	var users []models.User

	chain := database.Database.Table("users").
		Select("users.*", "user_roles.access_type").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Where("user_roles.customer_id = ?", currentCustomer.Id).
		Where("users.deleted_at IS NULL").
		Where("users.id != ?", user.Id)

	if len(searchQuery) > 0 {
		chain = chain.Where("users.name ILIKE ?", "%"+searchQuery+"%")
	}

	if err = chain.
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Scan(&users).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error fetching users.")
	}

	var total int64
	if err = database.Database.
		Table("users").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Where("user_roles.customer_id = ?", currentCustomer.Id).
		Where("users.deleted_at IS NULL").
		Where("users.id != ?", user.Id).
		Count(&total).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error counting users.")
	}

	var payload = make([]fiber.Map, 0)
	for _, user := range users {
		payload = append(payload, user.Json())
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
