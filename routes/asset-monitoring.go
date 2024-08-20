package routes

import (
	"github.com/gofiber/fiber/v2"
	assetMonitoringHandlers "github.com/predictive-edge-india/ehm-go/handlers/assetMonitoring"
	"github.com/predictive-edge-india/ehm-go/middlewares"
)

func AssetMonitoringRoutes(app fiber.Router) {
	group := app.Group("/asset-monitoring")
	group.Get("/:assetId/faults", middlewares.Protected(), assetMonitoringHandlers.FetchAssetFaults)
	group.Get("/:assetId/dgstatus", middlewares.Protected(), assetMonitoringHandlers.FetchAssetDGStatus)
	group.Get("/:assetId/last-location", middlewares.Protected(), assetMonitoringHandlers.FetchAssetLastLocation)
	group.Get("/:assetId/power-data", middlewares.Protected(), assetMonitoringHandlers.FetchPowerData)
}
