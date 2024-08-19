package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetStatusFlagForDevice(deviceId uuid.UUID) models.AlarmStatusFlag {
	var alarmStatusFlag models.AlarmStatusFlag
	Database.Where("device_id = ?", deviceId).Find(&alarmStatusFlag)
	return alarmStatusFlag
}

func GetStatusFlagForAsset(assetId uuid.UUID) models.AlarmStatusFlag {
	var alarmStatusFlag models.AlarmStatusFlag
	Database.Model(&models.AlarmStatusFlag{}).
		Joins("JOIN devices ON devices.id = alarm_status_flags.device_id").
		Joins("JOIN asset_devices ON asset_devices.device_id = devices.id").
		Joins("JOIN assets ON assets.id = asset_devices.asset_id").
		Where("assets.id = ?", assetId).
		Order("alarm_status_flags.created_at DESC").
		Find(&alarmStatusFlag)
	return alarmStatusFlag
}
