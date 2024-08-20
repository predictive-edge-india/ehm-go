package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetLastLocationForAsset(assetId uuid.UUID) models.DeviceLastLocation {
	var lastLocation models.DeviceLastLocation
	Database.Model(&models.DeviceLastLocation{}).
		Select("device_last_locations.id, ST_AsGeoJSON(device_last_locations.position) as position, device_last_locations.created_at").
		Joins("JOIN devices ON devices.id = device_last_locations.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("device_last_locations.created_at DESC").
		Find(&lastLocation)
	return lastLocation
}
