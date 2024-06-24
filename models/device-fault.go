package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceFault struct {
	gorm.Model
	Id                      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	LowOilPressure          bool      `gorm:"column:low_oil_pressure" json:"lowOilPressure"`
	HighCoolantTemperature  bool      `gorm:"column:high_coolant_temperature" json:"highCoolantTemperature"`
	LowFuelLevel            bool      `gorm:"column:low_fuel_level" json:"lowFuelLevel"`
	WaterLevelSwitch        bool      `gorm:"column:water_level_switch" json:"waterLevelSwitch"`
	UnderSpeed              bool      `gorm:"column:under_speed" json:"underSpeed"`
	OverSpeed               bool      `gorm:"column:over_speed" json:"overSpeed"`
	FailToStart             bool      `gorm:"column:fail_to_start" json:"failToStart"`
	FailToStop              bool      `gorm:"column:fail_to_stop" json:"failToStop"`
	ReversePowerDetected    bool      `gorm:"column:reverse_power_detected" json:"reversePowerDetected"`
	LowLoadAlarm            bool      `gorm:"column:low_load_alarm" json:"lowLoadAlarm"`
	GeneratorLowFrequency   bool      `gorm:"column:generator_low_frequency" json:"generatorLowFrequency"`
	GeneratorHighFrequency  bool      `gorm:"column:generator_high_frequency" json:"generatorHighFrequency"`
	GeneratorHighCurrent    bool      `gorm:"column:generator_high_current" json:"generatorHighCurrent"`
	GeneratorOverload       bool      `gorm:"column:generator_overload" json:"generatorOverload"`
	UnbalancedLoad          bool      `gorm:"column:unbalanced_load" json:"unbalancedLoad"`
	EmergencyStrip          bool      `gorm:"column:emergency_strip" json:"emergencyStrip"`
	ChargeAlternatorFailure bool      `gorm:"column:charge_alternator_failure" json:"chargeAlternatorFailure"`
	MaintenanceDue          bool      `gorm:"column:maintenance_due" json:"maintenanceDue"`
	AftActivationTimeout    bool      `gorm:"column:aft_activation_timeout" json:"aftActivationTimeout"`
	AshLoadMaintainance     bool      `gorm:"column:ash_load_maintainance" json:"ashLoadMaintainance"`
	BatteryLowVoltage       bool      `gorm:"column:battery_low_voltage" json:"batteryLowVoltage"`
	BatteryHighVoltage      bool      `gorm:"column:battery_high_voltage" json:"batteryHighVoltage"`
	TemperatureCircuitOpen  bool      `gorm:"column:temperature_circuit_open" json:"temperatureCircuitOpen"`
	LopBatteryShort         bool      `gorm:"column:lop_battery_short" json:"lopBatteryShort"`
	FuelTheft               bool      `gorm:"column:fuel_theft" json:"fuelTheft"`
	MagneticPickUpFault     bool      `gorm:"column:magnetic_pick_up_fault" json:"magneticPickUpFault"`
	OilPressureOpenCircuit  bool      `gorm:"column:oil_pressure_open_circuit" json:"oilPressureOpenCircuit"`
	L1LowVoltage            bool      `gorm:"column:l1_low_voltage" json:"l1LowVoltage"`
	L1HighVoltage           bool      `gorm:"column:l1_high_voltage" json:"l1HighVoltage"`
	L2LowVoltage            bool      `gorm:"column:l2_low_voltage" json:"l2LowVoltage"`
	L2HighVoltage           bool      `gorm:"column:l2_high_voltage" json:"l2HighVoltage"`
	L3LowVoltage            bool      `gorm:"column:l3_low_voltage" json:"l3LowVoltage"`
	L3HighVoltage           bool      `gorm:"column:l3_high_voltage" json:"l3HighVoltage"`

	EhmDeviceId *uuid.UUID `gorm:"type:uuid;column:ehm_device_id" json:"ehmDeviceId"`
	EhmDevice   EhmDevice  `gorm:"foreignKey:EhmDeviceId"`
}

func (u DeviceFault) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":                      u.Id,
		"lowOilPressure":          u.LowOilPressure,
		"highCoolantTemperature":  u.HighCoolantTemperature,
		"lowFuelLevel":            u.LowFuelLevel,
		"waterLevelSwitch":        u.WaterLevelSwitch,
		"underSpeed":              u.UnderSpeed,
		"overSpeed":               u.OverSpeed,
		"failToStart":             u.FailToStart,
		"failToStop":              u.FailToStop,
		"reversePowerDetected":    u.ReversePowerDetected,
		"lowLoadAlarm":            u.LowLoadAlarm,
		"generatorLowFrequency":   u.GeneratorLowFrequency,
		"generatorHighFrequency":  u.GeneratorHighFrequency,
		"generatorHighCurrent":    u.GeneratorHighCurrent,
		"generatorOverload":       u.GeneratorOverload,
		"unbalancedLoad":          u.UnbalancedLoad,
		"emergencyStrip":          u.EmergencyStrip,
		"chargeAlternatorFailure": u.ChargeAlternatorFailure,
		"maintenanceDue":          u.MaintenanceDue,
		"aftActivationTimeout":    u.AftActivationTimeout,
		"ashLoadMaintainance":     u.AshLoadMaintainance,
		"batteryLowVoltage":       u.BatteryLowVoltage,
		"batteryHighVoltage":      u.BatteryHighVoltage,
		"temperatureCircuitOpen":  u.TemperatureCircuitOpen,
		"lopBatteryShort":         u.LopBatteryShort,
		"fuelTheft":               u.FuelTheft,
		"magneticPickUpFault":     u.MagneticPickUpFault,
		"oilPressureOpenCircuit":  u.OilPressureOpenCircuit,
		"l1LowVoltage":            u.L1LowVoltage,
		"l1HighVoltage":           u.L1HighVoltage,
		"l2LowVoltage":            u.L2LowVoltage,
		"l2HighVoltage":           u.L2HighVoltage,
		"l3LowVoltage":            u.L3LowVoltage,
		"l3HighVoltage":           u.L3HighVoltage,
		"createdAt":               u.CreatedAt,
	}

	return payload
}
