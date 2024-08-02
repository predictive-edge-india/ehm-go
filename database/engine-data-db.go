package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetEngineDataForDevice(deviceId uuid.UUID, paramType string) models.EngineParam {
	var engineData models.EngineParam
	Database.Order("created_at DESC").First(&engineData,
		"ehm_device_id = ?",
		deviceId,
	)
	return engineData
}
