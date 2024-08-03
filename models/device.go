package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	SerialNo string    `gorm:"column:serial_no" json:"serialNo"`

	DeviceTypeId uuid.UUID  `gorm:"column:device_type_id" json:"deviceTypeId"`
	DeviceType   DeviceType `gorm:"foreignKey:DeviceTypeId"`
}

func (u Device) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"serialNo":  u.SerialNo,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Device) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"serialNo":  u.SerialNo,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Device) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
