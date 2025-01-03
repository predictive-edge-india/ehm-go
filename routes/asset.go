package routes

import (
	"github.com/gofiber/fiber/v2"
	assetHandlers "github.com/predictive-edge-india/ehm-go/handlers/assets"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetRoutes(app fiber.Router) {
	group := app.Group("/assets")
	group.Get("/formdata", middlewares.Protected(), assetHandlers.FetchAssetFormData)
	group.Get("/devices/formdata", middlewares.Protected(), assetHandlers.FetchAssignAssetDeviceFormData)

	group.Delete("/:assetId/devices/:deviceId", middlewares.Protected(), assetHandlers.UnassignAssetDevice)
	group.Post("/:assetId/devices", middlewares.Protected(), assetHandlers.AssignAssetDevice)
	group.Get("/:assetId/devices", middlewares.Protected(), assetHandlers.FetchAssetDevices)
	group.Post("/:assetId/limits", middlewares.Protected(), assetHandlers.UpdatePowerLimit)
	group.Get("/:assetId/limits", middlewares.Protected(), assetHandlers.GetPowerLimit)
	group.Patch("/:assetId", middlewares.Protected(), assetHandlers.UpdateAssetDetails)
	group.Get("/:assetId", middlewares.Protected(), assetHandlers.FetchAssetDetails)

	group.Get("/", middlewares.Protected(), assetHandlers.FetchAssets)
	group.Post("/", middlewares.Protected(), assetHandlers.CreateNewAsset)
}
