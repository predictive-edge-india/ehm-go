package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const DefaultPage = 1
const DefaultPerPage = 30

func GetPagination(c *fiber.Ctx) (int, int) {
	page, err := strconv.Atoi(c.Query("page", strconv.Itoa(DefaultPage)))
	if err != nil {
		page = DefaultPage
	}
	perPage, err := strconv.Atoi(c.Query("per_page", strconv.Itoa(DefaultPerPage)))
	if err != nil {
		perPage = DefaultPerPage
	}

	return page, perPage
}
