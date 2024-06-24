package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EngineParam struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	CoolantTemperature uint8 `gorm:"column:coolant_temperature" json:"coolantTemperature"`
	FuelTemperature    uint8 `gorm:"column:fuel_temperature" json:"fuelTemperature"`
	StarterMode        uint8 `gorm:"column:starter_mode" json:"starterMode"`

	Rpm                        float32 `gorm:"column:rpm" json:"rpm"`
	EngineDemandPercentTorque  int8    `gorm:"column:engine_demand_percent_torque" json:"engineDemandPercentTorque"`
	ActualPercentTorqueHighRes float32 `gorm:"column:actual_percent_torque_high_res" json:"actualPercentTorqueHighRes"`
	DriverDemandPercentTorque  int8    `gorm:"column:driver_demand_percent_torque" json:"driverDemandPercentTorque"`
	ActualEnginePercentTorque  int8    `gorm:"column:actual_engine_percent_torque" json:"actualEnginePercentTorque"`
	TorqueMode                 uint8   `gorm:"column:torque_mode" json:"torqueMode"`

	HoursByEms float32 `gorm:"column:hours_by_ems" json:"hoursByEms"`
	Second     float32 `gorm:"column:second" json:"second"`
	Minute     uint8   `gorm:"column:minute" json:"minute"`
	Hour       uint8   `gorm:"column:hour" json:"hour"`
	Month      uint8   `gorm:"column:month" json:"month"`
	Day        uint8   `gorm:"column:day" json:"day"`
	Year       uint16  `gorm:"column:year" json:"year"`

	RequestedSpeed                 float32 `gorm:"column:requested_speed" json:"requestedSpeed"`
	RequestedTorque                int8    `gorm:"column:requested_torque" json:"requestedTorque"`
	RequestedTorqueHighRes         float32 `gorm:"column:requested_torque_high_res" json:"requestedTorqueHighRes"`
	WaterInFuel                    uint8   `gorm:"column:water_in_fuel" json:"waterInFuel"`
	InducementLevel                uint8   `gorm:"column:inducement_level" json:"inducementLevel"`
	FuelEconomy                    float32 `gorm:"column:fuel_economy" json:"fuelEconomy"`
	FuelConsumptionTotalHighRes    float32 `gorm:"column:fuel_consumption_total_high_res" json:"fuelConsumptionTotalHighRes"`
	BatteryPotential               float32 `gorm:"column:battery_potential" json:"batteryPotential"`
	MalfunctionLampIndicatorStatus uint8   `gorm:"column:malfunction_lamp_indicator_status" json:"malfunctionLampIndicatorStatus"`
	StopLampStatus                 uint8   `gorm:"column:stop_lamp_status" json:"stopLampStatus"`
	WarningLampStatus              uint8   `gorm:"column:warning_lamp_status" json:"warningLampStatus"`
	UreaPercentageInWater          float32 `gorm:"column:urea_percentage_in_water" json:"ureaPercentageInWater"`
	AdblueTankLevel                float32 `gorm:"column:adblue_tank_level" json:"adblueTankLevel"`
	ParkBrakeStatus                uint8   `gorm:"column:park_brake_status" json:"parkBrakeStatus"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u EngineParam) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":                             u.Id,
		"coolantTemperature":             u.CoolantTemperature,
		"fuelTemperature":                u.FuelTemperature,
		"starterMode":                    u.StarterMode,
		"rpm":                            u.Rpm,
		"engineDemandPercentTorque":      u.EngineDemandPercentTorque,
		"actualPercentTorqueHighRes":     u.ActualPercentTorqueHighRes,
		"driverDemandPercentTorque":      u.DriverDemandPercentTorque,
		"actualEnginePercentTorque":      u.ActualEnginePercentTorque,
		"torqueMode":                     u.TorqueMode,
		"hoursByEms":                     u.HoursByEms,
		"second":                         u.Second,
		"minute":                         u.Minute,
		"hour":                           u.Hour,
		"month":                          u.Month,
		"day":                            u.Day,
		"year":                           u.Year,
		"requestedSpeed":                 u.RequestedSpeed,
		"requestedTorque":                u.RequestedTorque,
		"requestedTorqueHighRes":         u.RequestedTorqueHighRes,
		"waterInFuel":                    u.WaterInFuel,
		"inducementLevel":                u.InducementLevel,
		"fuelEconomy":                    u.FuelEconomy,
		"fuelConsumptionTotalHighRes":    u.FuelConsumptionTotalHighRes,
		"batteryPotential":               u.BatteryPotential,
		"malfunctionLampIndicatorStatus": u.MalfunctionLampIndicatorStatus,
		"stopLampStatus":                 u.StopLampStatus,
		"warningLampStatus":              u.WarningLampStatus,
		"ureaPercentageInWater":          u.UreaPercentageInWater,
		"adblueTankLevel":                u.AdblueTankLevel,
		"parkBrakeStatus":                u.ParkBrakeStatus,
		"createdAt":                      u.CreatedAt,
	}

	return payload
}
