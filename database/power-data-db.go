package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetPowerDataForAsset(assetId uuid.UUID) models.PowerData {
	var powerData models.PowerData
	Database.Model(&models.PowerData{}).
		Joins("JOIN devices ON devices.id = power_data.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("power_data.created_at DESC").
		Find(&powerData)
	return powerData
}
