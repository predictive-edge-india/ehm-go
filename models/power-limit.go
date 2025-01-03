package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PowerLimit struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Type       string    `gorm:"column:type" json:"type"`
	UpperLimit int16     `gorm:"column:upper_limit" json:"upperLimit"`
	LowerLimit int16     `gorm:"column:lower_limit" json:"lowerLimit"`

	AssetId *uuid.UUID `gorm:"column:asset_id" json:"assetId"`
	Asset   Asset      `gorm:"foreignKey:AssetId"`
}

func (u PowerLimit) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"type":       u.Type,
		"upperLimit": u.UpperLimit,
		"lowerLimit": u.LowerLimit,
		"createdAt":  u.CreatedAt,
	}

	if u.AssetId != nil {
		payload["asset"] = map[string]interface{}{
			"id":   u.Asset.Id,
			"name": u.Asset.Name,
		}
	}
	return payload
}

func (u PowerLimit) IsIdNull() bool {
	return u.Id == uuid.Nil
}
