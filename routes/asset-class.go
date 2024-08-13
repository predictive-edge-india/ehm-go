package routes

import (
	"github.com/gofiber/fiber/v2"
	assetClassHandlers "github.com/predictive-edge-india/ehm-go/handlers/assetClasses"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetClassRoutes(app fiber.Router) {
	group := app.Group("/asset-classes")
	group.Get("/:assetClassId", middlewares.Protected(), assetClassHandlers.FetchAssetClassDetails)
	group.Patch("/:assetClassId", middlewares.Protected(), assetClassHandlers.UpdateAssetClass)
	group.Delete("/:assetClassId", middlewares.Protected(), assetClassHandlers.DeleteAssetClass)
	group.Get("/", middlewares.Protected(), assetClassHandlers.FetchAssetClasses)
	group.Post("/", middlewares.Protected(), assetClassHandlers.CreateNewAssetClass)
}
