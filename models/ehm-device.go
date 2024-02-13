package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EhmDevice struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	SerialNo string    `gorm:"column:serial_no;not null" json:"serialNo"`
}

func (u EhmDevice) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"serialNo":  u.SerialNo,
		"createdAt": u.CreatedAt,
	}

	return payload
}
