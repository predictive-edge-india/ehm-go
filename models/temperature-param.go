package models

import (
	"github.com/google/uuid"
	"github.com/iisc/demo-go/helpers"
	"gorm.io/gorm"
)

type TemperatureParam struct {
	gorm.Model
	Id           uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Temperatures Float4Array `gorm:"type:double precision[];column:temperatures" json:"temperatures"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u TemperatureParam) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"createdAt": u.CreatedAt,
	}

	var temps []float32

	for _, v := range u.Temperatures {
		temps = append(temps, float32(helpers.ToFixed(float64(helpers.KelvinToCelsius(v)), 2)))
	}

	payload["temperatures"] = temps

	return payload
}
