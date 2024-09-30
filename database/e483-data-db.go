package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetE483DataForAsset(assetId uuid.UUID) models.E483CanData {
	var canData models.E483CanData
	Database.Model(&models.E483CanData{}).
		Joins("JOIN devices ON devices.id = e483_can_data.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("e483_can_data.created_at DESC").
		Find(&canData)
	return canData
}
