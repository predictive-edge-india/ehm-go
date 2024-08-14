package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetParameter struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Type       string    `gorm:"column:type" json:"type"`
	ParamOrder int16     `gorm:"column:param_order" json:"paramOrder"`

	AssetClassId *uuid.UUID `gorm:"column:asset_class_id" json:"assetClassId"`
	AssetClass   AssetClass `gorm:"foreignKey:AssetClassId"`
}

func (u AssetParameter) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"type":       u.Type,
		"paramOrder": u.ParamOrder,
		"createdAt":  u.CreatedAt,
	}

	if u.AssetClassId != nil {
		payload["assetClass"] = map[string]interface{}{
			"id":   u.AssetClass.Id,
			"name": u.AssetClass.Name,
		}
	}
	return payload
}
