package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"column:name" json:"name"`

	AssetClassId *uuid.UUID `gorm:"column:asset_class_id" json:"assetClassId"`
	AssetClass   AssetClass `gorm:"foreignKey:AssetClassId"`

	CustomerId *uuid.UUID `gorm:"column:customer_id" json:"customerId"`
	Customer   Customer   `gorm:"foreignKey:CustomerId"`
}

func (u Asset) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Asset) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Asset) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
