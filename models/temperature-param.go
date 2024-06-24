package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemperatureParam struct {
	gorm.Model
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Temperature []float32 `gorm:"type:double precision[];column:temperature" json:"temperature"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u TemperatureParam) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":          u.Id,
		"temperature": u.Temperature,
		"createdAt":   u.CreatedAt,
	}

	return payload
}
