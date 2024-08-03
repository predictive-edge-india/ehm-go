package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetDevice struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`
	AssetId  uuid.UUID `gorm:"column:asset_id" json:"assetId"`
	Asset    Asset     `gorm:"foreignKey:AssetId"`
}

func (AssetDevice) TableName() string {
	return "asset_devices"
}

func (o AssetDevice) IsIdNull() bool {
	return o.Id.String() == uuid.Nil.String()
}
