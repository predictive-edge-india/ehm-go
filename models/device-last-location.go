package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceLastLocation struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Position GeoJson   `gorm:"column:position;" json:"position"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`
}

func (DeviceLastLocation) TableName() string {
	return "device_last_locations"
}
