package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CurrentParameter struct {
	gorm.Model
	Id          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	RMS         float64     `gorm:"column:rms;not null" json:"rms"`
	FFT         [][]float64 `gorm:"column:fft;not null;serializer:json" json:"fft"`
	ParamType   string      `gorm:"column:param_type;not null" json:"paramType"`
	PacketType  int32       `gorm:"-"`
	EhmDeviceId *uuid.UUID  `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice   `gorm:"foreignKey:EhmDeviceId"`
}

func (u CurrentParameter) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":   u.Id,
		"rms":  u.RMS,
		"fft":  u.FFT,
		"type": u.ParamType,
	}

	return payload
}
