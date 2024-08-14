package routes

import (
	"github.com/gofiber/fiber/v2"
	assetParameterHandlers "github.com/predictive-edge-india/ehm-go/handlers/assetParameters"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetParameterRoutes(app fiber.Router) {
	group := app.Group("/asset-parameters")
	group.Post("/", middlewares.Protected(), assetParameterHandlers.CreateNewAssetParameter)
	group.Get("/", middlewares.Protected(), assetParameterHandlers.FetchAssetParameters)
}
