package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindAssetDevice(assetId, deviceId uuid.UUID) models.AssetDevice {
	var assetDevice models.AssetDevice
	Database.Where("device_id = ? AND asset_id = ?", deviceId, assetId).Find(&assetDevice)
	return assetDevice
}
