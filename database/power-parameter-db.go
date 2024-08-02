package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetPowerParameterForDevice(deviceId uuid.UUID, paramType string) models.PowerParam {
	var powerParam models.PowerParam
	Database.Order("created_at DESC").First(&powerParam,
		"ehm_device_id = ?",
		deviceId,
	)
	return powerParam
}
