package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FuelPercentage struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Percentage int16     `gorm:"column:percentage" json:"percentage"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u FuelPercentage) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"percentage": u.Percentage,
		"createdAt":  u.CreatedAt,
	}

	return payload
}
