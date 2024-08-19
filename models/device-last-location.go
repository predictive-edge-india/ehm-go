package models

import (
	"fmt"

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

func (u DeviceLastLocation) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"position":  u.Position,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u DeviceLastLocation) MqttPayload() string {
	payload := fmt.Sprintf("%f,%f", u.Position.Coordinates[0], u.Position.Coordinates[1])
	return payload
}
