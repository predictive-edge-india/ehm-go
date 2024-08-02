package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func CreateCurrentParameter(currentParameter models.CurrentParameter) *models.CurrentParameter {
	if err := Database.Save(&currentParameter).Error; err != nil {
		log.Error().AnErr("CreateCurrentParameter", err).Send()
	}
	return &currentParameter
}

func FindLatestEhmDeviceCurrentParameter(deviceId uuid.UUID, paramType string) models.CurrentParameter {
	var currentParameter models.CurrentParameter
	Database.Order("created_at DESC").First(&currentParameter,
		"ehm_device_id = ? AND param_type = ?",
		deviceId,
		paramType,
	)
	return currentParameter
}
