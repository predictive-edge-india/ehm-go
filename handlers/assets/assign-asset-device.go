package assetHandlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

func AssignAssetDevice(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	assetIdStr := c.Params("assetId")
	assetId, err := uuid.Parse(assetIdStr)
	if err != nil {
		return helpers.BadRequestError(c, "Invalid UUID!")
	}

	asset := database.FindAssetById(assetId)
	if asset.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Asset")
	}

	jsonBody := struct {
		Device string `json:"device" validate:"required,uuid4"`
	}{}

	if err := c.BodyParser(&jsonBody); err != nil {
		return helpers.BadRequestError(c, "Error parsing body!")
	}

	validate := validator.New()
	if err := validate.Struct(jsonBody); err != nil {
		return helpers.BadRequestError(c, "Please check your request!")
	}

	device := database.FindUnassignedDeviceById(uuid.MustParse(jsonBody.Device))
	if device.IsIdNull() {
		return helpers.ResourceNotFoundError(c, "Device")
	}

	assetDevice := models.AssetDevice{
		AssetId:  asset.Id,
		DeviceId: device.Id,
	}

	if err := database.Database.Create(&assetDevice).Error; err != nil {
		return helpers.BadRequestError(c, "There was an error assigning device!")
	}

	return c.JSON(helpers.BuildResponse("Success"))
}
