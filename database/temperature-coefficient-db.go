package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetTemperatureCoefficientForDevice(deviceId uuid.UUID) models.TemperatureParam {
	var tempParam models.TemperatureParam
	Database.Order("created_at DESC").First(&tempParam,
		"ehm_device_id = ?",
		deviceId,
	)
	return tempParam
}
