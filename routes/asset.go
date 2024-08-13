package routes

import (
	"github.com/gofiber/fiber/v2"
	assetHandlers "github.com/predictive-edge-india/ehm-go/handlers/assets"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetRoutes(app fiber.Router) {
	group := app.Group("/assets")
	group.Get("/formdata", middlewares.Protected(), assetHandlers.FetchAssetFormData)
	group.Get("/", middlewares.Protected(), assetHandlers.FetchAssets)
	group.Post("/", middlewares.Protected(), assetHandlers.CreateNewAsset)
}
