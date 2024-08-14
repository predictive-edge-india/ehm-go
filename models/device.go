package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	SerialNo string    `gorm:"column:serial_no" json:"serialNo"`

	Imei1  string         `gorm:"column:imei1" json:"imei1"`
	Imei2  sql.NullString `gorm:"column:imei2" json:"imei2"`
	Phone1 string         `gorm:"column:phone1" json:"phone1"`
	Phone2 sql.NullString `gorm:"column:phone2" json:"phone2"`

	Note sql.NullString `gorm:"column:note" json:"note"`

	DeviceTypeId uuid.UUID  `gorm:"column:device_type_id" json:"deviceTypeId"`
	DeviceType   DeviceType `gorm:"foreignKey:DeviceTypeId"`
}

func (u Device) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"serialNo":  u.SerialNo,
		"imei1":     u.Imei1,
		"phone1":    u.Phone1,
		"createdAt": u.CreatedAt,
	}

	if u.Imei2.Valid {
		payload["imei2"] = u.Imei2.String
	}
	if u.Phone2.Valid {
		payload["phone2"] = u.Phone2.String
	}
	if u.Note.Valid {
		payload["note"] = u.Note.String
	}

	if u.DeviceTypeId != uuid.Nil {
		payload["deviceType"] = u.DeviceType.ShortJson()
	}

	return payload
}

func (u Device) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"serialNo":  u.SerialNo,
		"phone1":    u.Phone1,
		"imei1":     u.Imei1,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Device) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
