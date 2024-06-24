package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PowerParam struct {
	gorm.Model
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	L1NVoltage     uint16    `gorm:"column:l1_n_voltage" json:"l1NVoltage"`
	L2NVoltage     uint16    `gorm:"column:l2_n_voltage" json:"l2NVoltage"`
	L3NVoltage     uint16    `gorm:"column:l3_n_voltage" json:"l3NVoltage"`
	L1L2Voltage    uint16    `gorm:"column:l1_l2_voltage" json:"l1L2Voltage"`
	L2L3Voltage    uint16    `gorm:"column:l2_l3_voltage" json:"l2L3Voltage"`
	L2L1Voltage    uint16    `gorm:"column:l2_l1_voltage" json:"l2L1Voltage"`
	L1Current      uint16    `gorm:"column:l1_current" json:"l1Current"`
	L2Current      uint16    `gorm:"column:l2_current" json:"l2Current"`
	L3Current      uint16    `gorm:"column:l3_current" json:"l3Current"`
	TotalWatts     uint16    `gorm:"column:total_watts" json:"totalWatts"`
	LoadPercentage uint16    `gorm:"column:load_percentage" json:"loadPercentage"`
	L1VA           uint16    `gorm:"column:l1_va" json:"l1VA"`
	L2VA           uint16    `gorm:"column:l2_va" json:"l2VA"`
	L3VA           uint16    `gorm:"column:l3_va" json:"l3VA"`
	TotalVA        uint16    `gorm:"column:total_va" json:"totalVA"`

	RFrequency uint8 `gorm:"column:r_frequency" json:"rFrequency"`
	YFrequency uint8 `gorm:"column:y_frequency" json:"yFrequency"`
	BFrequency uint8 `gorm:"column:b_frequency" json:"bFrequency"`

	PfL1  float32 `gorm:"column:pf_l1" json:"pfL1"`
	PfL2  float32 `gorm:"column:pf_l2" json:"pfL2"`
	PfL3  float32 `gorm:"column:pf_l3" json:"pfL3"`
	PfAvg float32 `gorm:"column:pf_avg" json:"pfAvg"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u PowerParam) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":             u.Id,
		"l1NVoltage":     u.L1NVoltage,
		"l2NVoltage":     u.L2NVoltage,
		"l3NVoltage":     u.L3NVoltage,
		"l1L2Voltage":    u.L1L2Voltage,
		"l2L3Voltage":    u.L2L3Voltage,
		"l2L1Voltage":    u.L2L1Voltage,
		"l1Current":      u.L1Current,
		"l2Current":      u.L2Current,
		"l3Current":      u.L3Current,
		"totalWatts":     u.TotalWatts,
		"loadPercentage": u.LoadPercentage,
		"l1VA":           u.L1VA,
		"l2VA":           u.L2VA,
		"l3VA":           u.L3VA,
		"totalVA":        u.TotalVA,
		"rFrequency":     u.RFrequency,
		"yFrequency":     u.YFrequency,
		"bFrequency":     u.BFrequency,
		"pfL1":           u.PfL1,
		"pfL2":           u.PfL2,
		"pfL3":           u.PfL3,
		"pfAvg":          u.PfAvg,
		"createdAt":      u.CreatedAt,
	}

	return payload
}
