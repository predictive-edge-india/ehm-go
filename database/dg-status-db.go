package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetDGStatusForAsset(assetId uuid.UUID) models.DGStatus {
	var dgStatus models.DGStatus
	Database.Model(&models.DGStatus{}).
		Joins("JOIN devices ON devices.id = dg_statuses.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("dg_statuses.created_at DESC").
		Find(&dgStatus)
	return dgStatus
}
