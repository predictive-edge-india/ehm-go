package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceType struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"column:name" json:"name"`
}

func (u DeviceType) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u DeviceType) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u DeviceType) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
