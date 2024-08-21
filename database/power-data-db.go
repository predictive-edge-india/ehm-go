package database

import (
	"fmt"
	"time"

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

type PowerDataParam struct {
	Value string `gorm:"column:value"`
}

type PowerDataParamWithTime struct {
	Value string `gorm:"column:value"`
	Time  string `gorm:"column:created_at"`
}

func GetPowerParameterForAsset(assetId uuid.UUID, paramKey string) PowerDataParam {
	var powerData PowerDataParam
	Database.Model(&models.PowerData{}).
		Select(fmt.Sprintf("CAST(power_data.%s as text) as value", paramKey)).
		Joins("JOIN devices ON devices.id = power_data.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("power_data.created_at DESC").
		Find(&powerData)
	return powerData
}

func GetPowerParameterForAssetWithRange(assetId uuid.UUID, paramKey string, startDate, endDate time.Time) []PowerDataParamWithTime {
	var powerData []PowerDataParamWithTime
	Database.Model(&models.PowerData{}).
		Select(fmt.Sprintf("CAST(power_data.%s as text) as value, power_data.created_at", paramKey)).
		Joins("JOIN devices ON devices.id = power_data.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Where("power_data.created_at BETWEEN ? AND ?", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).
		Order("power_data.created_at DESC").
		Find(&powerData)
	return powerData
}
