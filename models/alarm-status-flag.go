package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlarmStatusFlag struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Statuses []uint8   `gorm:"column:statuses" json:"statuses"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`
}

func (AlarmStatusFlag) TableName() string {
	return "alarm_status_flags"
}

func (flag *AlarmStatusFlag) MqttPayload() string {
	return fmt.Sprintf("%d", flag.Statuses)
}
