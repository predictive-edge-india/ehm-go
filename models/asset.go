package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Id   uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`

	Make      string `gorm:"column:make" json:"make"`
	ModelName string `gorm:"column:model" json:"model"`

	AssetClassId *uuid.UUID `gorm:"column:asset_class_id" json:"assetClassId"`
	AssetClass   AssetClass `gorm:"foreignKey:AssetClassId"`

	CustomerId *uuid.UUID `gorm:"column:customer_id" json:"customerId"`
	Customer   Customer   `gorm:"foreignKey:CustomerId"`
}

func (u Asset) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"make":      u.Make,
		"model":     u.ModelName,
		"createdAt": u.CreatedAt,
	}

	if u.Name.Valid {
		payload["name"] = u.Name.String
	}

	if u.AssetClassId != nil {
		payload["assetClass"] = map[string]interface{}{
			"id":   u.AssetClass.Id,
			"name": u.AssetClass.Name,
		}
	}
	if u.CustomerId != nil {
		payload["customer"] = map[string]interface{}{
			"id":   u.Customer.Id,
			"name": u.Customer.Name,
		}
	}
	return payload
}

func (u Asset) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"make":      u.Make,
		"model":     u.ModelName,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Asset) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
