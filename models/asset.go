package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"column:name" json:"name"`

	Make      string `gorm:"column:make" json:"make"`
	ModelName string `gorm:"column:model" json:"model"`

	TankCapacity sql.NullInt32 `gorm:"column:tank_capacity" json:"tankCapacity"`

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
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}

	if !u.AssetClass.IsIdNull() {
		payload["assetClass"] = map[string]interface{}{
			"id":   u.AssetClass.Id,
			"name": u.AssetClass.Name,
		}
	}
	if !u.Customer.IsIdNull() {
		payload["customer"] = map[string]interface{}{
			"id":   u.Customer.Id,
			"name": u.Customer.Name,
		}
	}

	if u.TankCapacity.Valid {
		payload["tankCapacity"] = u.TankCapacity.Int32
	}
	return payload
}

func (u Asset) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"make":      u.Make,
		"model":     u.ModelName,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Asset) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
