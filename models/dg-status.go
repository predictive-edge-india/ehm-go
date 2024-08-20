package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DGStatus struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`

	SGCMonitoringMode       bool   `gorm:"column:sgc_monitoring_mode" json:"sgcMonitoringMode"`
	MainsHealthy            bool   `gorm:"column:mains_healthy" json:"mainsHealthy"`
	DGOperationMode         uint8  `gorm:"column:dg_operation_mode" json:"dgOperationMode"`
	LoadOnMains             bool   `gorm:"column:load_on_mains" json:"loadOnMains"`
	LoadOnDG                bool   `gorm:"column:load_on_dg" json:"loadOnDG"`
	CurrentDGStatus         bool   `gorm:"column:current_dg_status" json:"currentDGStatus"`
	DGStoppedNormally       bool   `gorm:"column:dg_stopped_normally" json:"dgStoppedNormally"`
	DGStoppedWithFault      bool   `gorm:"column:dg_stopped_with_fault" json:"dgStoppedWithFault"`
	DGFailToStart           bool   `gorm:"column:dg_fail_to_start" json:"dgFailToStart"`
	GenAvailable            bool   `gorm:"column:gen_available" json:"genAvailable"`
	CommonShutDown          bool   `gorm:"column:common_shutdown" json:"commonShutDown"`
	CommonElectricTrip      bool   `gorm:"column:common_electric_trip" json:"commonElectricTrip"`
	CommonWarning           bool   `gorm:"column:common_warning" json:"commonWarning"`
	CommonNotification      bool   `gorm:"column:common_notification" json:"commonNotification"`
	CurrentTimeStampMin     uint8  `gorm:"column:current_timestamp_min" json:"currentTimeStampMin"`
	CurrentTimeStampSec     uint8  `gorm:"column:current_timestamp_sec" json:"currentTimeStampSec"`
	CurrentTimeStampWeekDay uint8  `gorm:"column:current_timestamp_week_day" json:"currentTimeStampWeekDay"`
	CurrentTimeStampHour    uint8  `gorm:"column:current_timestamp_hour" json:"currentTimeStampHour"`
	CurrentTimeStampMonth   uint8  `gorm:"column:current_timestamp_month" json:"currentTimeStampMonth"`
	CurrentTimeStampDay     uint8  `gorm:"column:current_timestamp_day" json:"currentTimeStampDay"`
	CurrentTimeStampYear    uint16 `gorm:"column:current_timestamp_year" json:"currentTimeStampYear"`
}

// 0,0,4,0,1,1,0,0,0,1,0,0,0,0,14,21,3,15,8,14,2024

func (d DGStatus) TableName() string {
	return "dg_statuses"
}

func (dg *DGStatus) GetUnixTimestamp() int64 {
	// Create a time.Time object from the timestamp fields
	timestamp := time.Date(
		int(dg.CurrentTimeStampYear),
		time.Month(dg.CurrentTimeStampMonth),
		int(dg.CurrentTimeStampDay),
		int(dg.CurrentTimeStampHour),
		int(dg.CurrentTimeStampMin),
		int(dg.CurrentTimeStampSec),
		0,        // Nanoseconds
		time.UTC, // Assuming UTC timezone
	)

	// Return the Unix timestamp
	return timestamp.Unix()
}

func (dg *DGStatus) MqttPayload() string {
	// Create a slice of string representations of each field
	fields := []string{
		fmt.Sprintf("%t", dg.SGCMonitoringMode),
		fmt.Sprintf("%t", dg.MainsHealthy),
		fmt.Sprintf("%d", dg.DGOperationMode),
		fmt.Sprintf("%t", dg.LoadOnMains),
		fmt.Sprintf("%t", dg.LoadOnDG),
		fmt.Sprintf("%t", dg.CurrentDGStatus),
		fmt.Sprintf("%t", dg.DGStoppedNormally),
		fmt.Sprintf("%t", dg.DGStoppedWithFault),
		fmt.Sprintf("%t", dg.DGFailToStart),
		fmt.Sprintf("%t", dg.GenAvailable),
		fmt.Sprintf("%t", dg.CommonShutDown),
		fmt.Sprintf("%t", dg.CommonElectricTrip),
		fmt.Sprintf("%t", dg.CommonWarning),
		fmt.Sprintf("%t", dg.CommonNotification),
		fmt.Sprintf("%d", dg.GetUnixTimestamp()),
	}

	// Join the fields with commas
	return strings.Join(fields, ",")
}

func (dg DGStatus) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":                 dg.Id,
		"sgcMonitoringMode":  dg.SGCMonitoringMode,
		"mainsHealthy":       dg.MainsHealthy,
		"dgOperationMode":    dg.DGOperationMode,
		"loadOnMains":        dg.LoadOnMains,
		"loadOnDG":           dg.LoadOnDG,
		"currentDGStatus":    dg.CurrentDGStatus,
		"dgStoppedNormally":  dg.DGStoppedNormally,
		"dgStoppedWithFault": dg.DGStoppedWithFault,
		"dgFailToStart":      dg.DGFailToStart,
		"genAvailable":       dg.GenAvailable,
		"commonShutDown":     dg.CommonShutDown,
		"commonElectricTrip": dg.CommonElectricTrip,
		"commonWarning":      dg.CommonWarning,
		"commonNotification": dg.CommonNotification,
		"timestamp":          dg.GetUnixTimestamp(),
	}
	return payload
}
