package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func GetFaultPanelForDevice(deviceId uuid.UUID, paramType string) models.FaultPanel {
	var faultPanel models.FaultPanel
	Database.Order("created_at DESC").First(&faultPanel,
		"ehm_device_id = ?",
		deviceId,
	)
	return faultPanel
}
