package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetClass struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"column:name" json:"name"`
}

func (u AssetClass) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u AssetClass) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u AssetClass) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
