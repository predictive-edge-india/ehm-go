package dashboardHandlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
)

type AssetLatestLocation struct {
	AssetID  uint           `json:"asset_id"`
	Name     string         `json:"name"`
	Position models.GeoJson `json:"position"`
	ReadAt   *time.Time     `json:"read_at"`
}

func FetchAssetLocations(c *fiber.Ctx) error {
	user := database.FindUserAuth(c)

	userRole, err := database.FindUserRoleForUser(c, user)
	if err != nil {
		return err
	}

	if userRole.AccessType != models.UserRoleEnum.SuperAdministrator.Number {
		return helpers.NotAuthorizedError(c, "You're not authorized!")
	}

	var assetLastLocations []AssetLatestLocation

	err = database.Database.Raw(`
			SELECT 
					assets.id, 
					assets.name, 
					ST_AsGeoJSON(device_last_locations.position) as position,
				device_last_locations.read_at
			FROM assets
			JOIN asset_devices ON assets.id = asset_devices.asset_id
			JOIN devices ON asset_devices.device_id = devices.id
			JOIN (
					SELECT DISTINCT ON (device_id) *
					FROM device_last_locations
					ORDER BY device_id, read_at DESC
			) AS device_last_locations ON devices.id = device_last_locations.device_id;
 		`).
		Scan(&assetLastLocations).Error
	if err != nil {
		return helpers.BadRequestError(c, "There was an error fetching assets.")
	}

	var assetJson []fiber.Map
	for _, asset := range assetLastLocations {
		obj := fiber.Map{
			"id":       asset.AssetID,
			"name":     asset.Name,
			"position": asset.Position,
		}
		if asset.ReadAt != nil {
			obj["readAt"] = asset.ReadAt
		}
		assetJson = append(assetJson, obj)
	}

	payload := fiber.Map{
		"assets": assetJson,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
