package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Id       uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string         `gorm:"column:name" json:"name"`
	LogoUrl  sql.NullString `gorm:"column:logo_url" json:"logoUrl"`
	Position GeoJson        `gorm:"column:position;" json:"position"`
	Email    string         `gorm:"column:email;" json:"email"`
	Phone    string         `gorm:"column:phone;" json:"phone"`

	Address1   string         `gorm:"column:address_1" json:"address1"`
	Address2   sql.NullString `gorm:"column:address_2" json:"address2"`
	City       string         `gorm:"column:city" json:"city"`
	State      string         `gorm:"column:state" json:"state"`
	Country    string         `gorm:"column:country" json:"country"`
	PostalCode int32          `gorm:"column:postal_code" json:"postalCode"`

	Gstin sql.NullString `gorm:"column:gstin" json:"gstin"`
}

func (u Customer) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"position":   u.Position,
		"email":      u.Email,
		"phone":      u.Phone,
		"address1":   u.Address1,
		"city":       u.City,
		"state":      u.State,
		"country":    u.Country,
		"postalCode": u.PostalCode,
		"createdAt":  u.CreatedAt,
	}
	if u.LogoUrl.Valid {
		payload["logoUrl"] = u.LogoUrl.String
	}
	if u.Address2.Valid {
		payload["address2"] = u.Address2.String
	}
	if u.Gstin.Valid {
		payload["gstin"] = u.Gstin.String
	}

	return payload
}

func (u Customer) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
	if u.LogoUrl.Valid {
		payload["logoUrl"] = u.LogoUrl.String
	}
	return payload
}

func (u Customer) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
