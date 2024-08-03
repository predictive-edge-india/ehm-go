package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	Id         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	AccessType int16      `gorm:"column:access_type" json:"accessType"`
	UserId     uuid.UUID  `gorm:"column:user_id" json:"userId"`
	User       User       `gorm:"foreignKey:UserId"`
	CustomerId *uuid.UUID `gorm:"column:customer_id" json:"customerId"`
	Customer   Customer   `gorm:"foreignKey:CustomerId"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

func (o UserRole) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         o.Id,
		"accessType": o.AccessType,
	}

	return payload
}
