package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type E483CanData struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`

	CoolantTemperature          uint8   `gorm:"column:coolant_temperature"`
	FuelTemperature             uint8   `gorm:"column:fuel_temperature"`
	StarterMode                 uint8   `gorm:"column:starter_mode"`
	RPM                         float32 `gorm:"column:rpm"`
	EngineDemandPercentTorque   int8    `gorm:"column:engine_demand_percent_torque"`
	ActualPercentTorqueHighRes  float32 `gorm:"column:actual_percent_torque_high_res"`
	DriverDemandPercentTorque   int8    `gorm:"column:driver_demand_percent_torque"`
	ActualEnginePercentTorque   int8    `gorm:"column:actual_engine_percent_torque"`
	TorqueMode                  uint8   `gorm:"column:torque_mode"`
	HoursByEMS                  float64 `gorm:"column:hours_by_ems"`
	WaterInFuel                 uint8   `gorm:"column:water_in_fuel"`
	FuelEconomy                 float32 `gorm:"column:fuel_economy"`
	FuelConsumptionTotalHighRes float64 `gorm:"column:fuel_consumption_total_high_res"`
	BatteryPotential            float32 `gorm:"column:battery_potential"`
	UreaPercentageInWater       float32 `gorm:"column:urea_percentage_in_water"`
	AdblueTankLevel             float32 `gorm:"column:adblue_tank_level"`
	EngineOilPressure           uint16  `gorm:"column:engine_oil_pressure"`
}

func (p E483CanData) Json() map[string]interface{} {
	return map[string]interface{}{
		"id":                          p.Id,
		"coolantTemperature":          p.CoolantTemperature,
		"fuelTemperature":             p.FuelTemperature,
		"starterMode":                 p.StarterMode,
		"rpm":                         p.RPM,
		"engineDemandPercentTorque":   p.EngineDemandPercentTorque,
		"actualPercentTorqueHighRes":  p.ActualPercentTorqueHighRes,
		"driverDemandPercentTorque":   p.DriverDemandPercentTorque,
		"actualEnginePercentTorque":   p.ActualEnginePercentTorque,
		"torqueMode":                  p.TorqueMode,
		"hoursByEMS":                  p.HoursByEMS,
		"waterInFuel":                 p.WaterInFuel,
		"fuelEconomy":                 p.FuelEconomy,
		"fuelConsumptionTotalHighRes": p.FuelConsumptionTotalHighRes,
		"batteryPotential":            p.BatteryPotential,
		"ureaPercentageInWater":       p.UreaPercentageInWater,
		"adblueTankLevel":             p.AdblueTankLevel,
		"engineOilPressure":           p.EngineOilPressure,
	}
}

func (p E483CanData) MqttPayload() string {
	fields := []string{
		fmt.Sprintf("%d", p.CoolantTemperature),
		fmt.Sprintf("%d", p.FuelTemperature),
		fmt.Sprintf("%d", p.StarterMode),
		fmt.Sprintf("%f", p.RPM),
		fmt.Sprintf("%d", p.EngineDemandPercentTorque),
		fmt.Sprintf("%f", p.ActualPercentTorqueHighRes),
		fmt.Sprintf("%d", p.DriverDemandPercentTorque),
		fmt.Sprintf("%d", p.ActualEnginePercentTorque),
		fmt.Sprintf("%d", p.TorqueMode),
		fmt.Sprintf("%f", p.HoursByEMS),
		fmt.Sprintf("%d", p.WaterInFuel),
		fmt.Sprintf("%f", p.FuelEconomy),
		fmt.Sprintf("%f", p.FuelConsumptionTotalHighRes),
		fmt.Sprintf("%f", p.BatteryPotential),
		fmt.Sprintf("%f", p.UreaPercentageInWater),
		fmt.Sprintf("%f", p.AdblueTankLevel),
		fmt.Sprintf("%d", p.EngineOilPressure),
	}

	return strings.Join(fields, ",")
}
