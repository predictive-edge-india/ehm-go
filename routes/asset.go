package routes

import (
	"github.com/gofiber/fiber/v2"
	assetHandlers "github.com/predictive-edge-india/ehm-go/handlers/assets"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetRoutes(app fiber.Router) {
	assets := app.Group("/assets")
	assets.Get("/", middlewares.Protected(), assetHandlers.FetchAssets)
	assets.Post("/", middlewares.Protected(), assetHandlers.CreateNewAsset)
}
