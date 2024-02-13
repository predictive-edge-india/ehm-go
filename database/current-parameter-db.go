package database

import (
	"github.com/google/uuid"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func CreateCurrentParameter(currentParameter models.CurrentParameter) *models.CurrentParameter {
	err := Database.Save(&currentParameter).Error
	if err != nil {
		log.Errorln(err.Error())
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
