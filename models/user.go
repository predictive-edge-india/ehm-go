package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name         string         `gorm:"column:name" json:"name"`
	Email        string         `gorm:"column:email" json:"email"`
	Phone        sql.NullString `gorm:"column:phone" json:"phone"`
	ProfilePic   sql.NullString `gorm:"column:profile_pic" json:"profilePic"`
	PasswordHash string         `gorm:"column:password_hash" json:"passwordHash"`
	Gender       sql.NullString `gorm:"column:gender" json:"gender"`
	BirthDate    sql.NullTime   `gorm:"column:birth_date" json:"birthDate"`
	UserRoles    []*UserRole
}

func (u User) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":    u.Id,
		"name":  u.Name,
		"phone": u.Phone,
	}
	return payload
}

func (u User) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"email":      u.Email,
		"phone":      nil,
		"profilePic": nil,
		"gender":     nil,
		"birthDate":  nil,
		"createdAt":  u.CreatedAt.Format(time.RFC3339),
	}
	if u.Phone.Valid {
		payload["phone"] = u.Phone.String
	}
	if u.ProfilePic.Valid {
		payload["profilePic"] = u.ProfilePic.String
	}
	if u.Gender.Valid {
		payload["gender"] = u.Gender.String
	}
	if u.BirthDate.Valid {
		payload["birthDate"] = u.BirthDate.Time.Format(time.RFC3339)
	}
	return payload
}

func (u User) IsIdNull() bool {
	return u.Id.String() == uuid.Nil.String()
}
