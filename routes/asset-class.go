package routes

import (
	"github.com/gofiber/fiber/v2"
	assetClassHandlers "github.com/predictive-edge-india/ehm-go/handlers/assetClass"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetClassRoutes(app fiber.Router) {
	assets := app.Group("/asset-classes")
	assets.Delete("/:assetClassId", middlewares.Protected(), assetClassHandlers.DeleteAssetClass)
	assets.Get("/", middlewares.Protected(), assetClassHandlers.FetchAssetClasses)
	assets.Post("/", middlewares.Protected(), assetClassHandlers.CreateNewAssetClass)
}
