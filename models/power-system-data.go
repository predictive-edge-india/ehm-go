package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PowerData struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	DeviceId uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Device   Device    `gorm:"foreignKey:DeviceId"`

	ProtocolRevision                  uint16 `gorm:"column:protocol_revision"`
	GeneratorL1NVoltage               uint16 `gorm:"column:generator_l1n_voltage"`
	GeneratorL2NVoltage               uint16 `gorm:"column:generator_l2n_voltage"`
	GeneratorL3NVoltage               uint16 `gorm:"column:generator_l3n_voltage"`
	GeneratorL1L2Voltage              uint16 `gorm:"column:generator_l1l2_voltage"`
	GeneratorL2L3Voltage              uint16 `gorm:"column:generator_l2l3_voltage"`
	GeneratorL3L1Voltage              uint16 `gorm:"column:generator_l3l1_voltage"`
	GeneratorRFrequency               uint16 `gorm:"column:generator_r_frequency"`
	GeneratorYFrequency               uint16 `gorm:"column:generator_y_frequency"`
	GeneratorBFrequency               uint16 `gorm:"column:generator_b_frequency"`
	GeneratorPowerFactorL1            int16  `gorm:"column:generator_power_factor_l1"`
	GeneratorPowerFactorL2            int16  `gorm:"column:generator_power_factor_l2"`
	GeneratorPowerFactorL3            int16  `gorm:"column:generator_power_factor_l3"`
	GeneratorAveragePowerFactor       int16  `gorm:"column:generator_average_power_factor"`
	MainsL1NVoltage                   uint16 `gorm:"column:mains_l1n_voltage"`
	MainsL2NVoltage                   uint16 `gorm:"column:mains_l2n_voltage"`
	MainsL3NVoltage                   uint16 `gorm:"column:mains_l3n_voltage"`
	MainsL1L2Voltage                  uint16 `gorm:"column:mains_l1l2_voltage"`
	MainsL2L3Voltage                  uint16 `gorm:"column:mains_l2l3_voltage"`
	MainsL3L1Voltage                  uint16 `gorm:"column:mains_l3l1_voltage"`
	MainsRFrequency                   uint16 `gorm:"column:mains_r_frequency"`
	MainsYFrequency                   uint16 `gorm:"column:mains_y_frequency"`
	MainsBFrequency                   uint16 `gorm:"column:mains_b_frequency"`
	LoadL1Current                     uint16 `gorm:"column:load_l1_current"`
	LoadL2Current                     uint16 `gorm:"column:load_l2_current"`
	LoadL3Current                     uint16 `gorm:"column:load_l3_current"`
	LoadL1Watts                       int16  `gorm:"column:load_l1_watts"`
	LoadL2Watts                       int16  `gorm:"column:load_l2_watts"`
	LoadL3Watts                       int16  `gorm:"column:load_l3_watts"`
	LoadTotalWatts                    int16  `gorm:"column:load_total_watts"`
	PercentageLoad                    int16  `gorm:"column:percentage_load"`
	LoadL1Va                          uint16 `gorm:"column:load_l1_va"`
	LoadL2Va                          uint16 `gorm:"column:load_l2_va"`
	LoadL3Va                          uint16 `gorm:"column:load_l3_va"`
	LoadTotalVa                       uint16 `gorm:"column:load_total_va"`
	LoadL1Var                         uint16 `gorm:"column:load_l1_var"`
	LoadL2Var                         uint16 `gorm:"column:load_l2_var"`
	LoadL3Var                         uint16 `gorm:"column:load_l3_var"`
	LoadTotalVar                      uint16 `gorm:"column:load_total_var"`
	GeneratorCumulativeEnergy         uint32 `gorm:"column:generator_cumulative_energy"`
	GeneratorCumulativeApparentEnergy uint32 `gorm:"column:generator_cumulative_apparent_energy"`
	GeneratorCumulativeReactiveEnergy uint32 `gorm:"column:generator_cumulative_reactive_energy"`
	MainsCumulativeEnergy             uint32 `gorm:"column:mains_cumulative_energy"`
	MainsCumulativeApparentEnergy     uint32 `gorm:"column:mains_cumulative_apparent_energy"`
	MainsCumulativeReactiveEnergy     uint32 `gorm:"column:mains_cumulative_reactive_energy"`
	OilPressure                       uint16 `gorm:"column:oil_pressure"`
	CoolantTemperature                int16  `gorm:"column:coolant_temperature"`
	FuelLevel                         uint16 `gorm:"column:fuel_level"`
	FuelLevelInLit                    uint16 `gorm:"column:fuel_level_in_lit"`
	ChargeAlternatorVoltage           uint16 `gorm:"column:charge_alternator_voltage"`
	BatteryVoltage                    uint16 `gorm:"column:battery_voltage"`
	EngineSpeed                       uint16 `gorm:"column:engine_speed"`
	NoOfStarts                        uint16 `gorm:"column:no_of_starts"`
	NoOfTrips                         uint16 `gorm:"column:no_of_trips"`
	EngRunHrs                         uint32 `gorm:"column:eng_run_hrs"`
	EngRunMin                         uint16 `gorm:"column:eng_run_min"`
	MainsRunHrs                       uint16 `gorm:"column:mains_run_hrs"`
	MainsRunMin                       uint16 `gorm:"column:mains_run_min"`
}

// Json returns a map of the PowerSystemData fields
func (p PowerData) Json() map[string]interface{} {
	return map[string]interface{}{
		"id":                                p.Id,
		"protocolRevision":                  p.ProtocolRevision,
		"generatorL1NVoltage":               p.GeneratorL1NVoltage,
		"generatorL2NVoltage":               p.GeneratorL2NVoltage,
		"generatorL3NVoltage":               p.GeneratorL3NVoltage,
		"generatorL1L2Voltage":              p.GeneratorL1L2Voltage,
		"generatorL2L3Voltage":              p.GeneratorL2L3Voltage,
		"generatorL3L1Voltage":              p.GeneratorL3L1Voltage,
		"generatorRFrequency":               p.GeneratorRFrequency,
		"generatorYFrequency":               p.GeneratorYFrequency,
		"generatorBFrequency":               p.GeneratorBFrequency,
		"generatorPowerFactorL1":            p.GeneratorPowerFactorL1,
		"generatorPowerFactorL2":            p.GeneratorPowerFactorL2,
		"generatorPowerFactorL3":            p.GeneratorPowerFactorL3,
		"generatorAveragePowerFactor":       p.GeneratorAveragePowerFactor,
		"mainsL1NVoltage":                   p.MainsL1NVoltage,
		"mainsL2NVoltage":                   p.MainsL2NVoltage,
		"mainsL3NVoltage":                   p.MainsL3NVoltage,
		"mainsL1L2Voltage":                  p.MainsL1L2Voltage,
		"mainsL2L3Voltage":                  p.MainsL2L3Voltage,
		"mainsL3L1Voltage":                  p.MainsL3L1Voltage,
		"mainsRFrequency":                   p.MainsRFrequency,
		"mainsYFrequency":                   p.MainsYFrequency,
		"mainsBFrequency":                   p.MainsBFrequency,
		"loadL1Current":                     p.LoadL1Current,
		"loadL2Current":                     p.LoadL2Current,
		"loadL3Current":                     p.LoadL3Current,
		"loadL1Watts":                       p.LoadL1Watts,
		"loadL2Watts":                       p.LoadL2Watts,
		"loadL3Watts":                       p.LoadL3Watts,
		"loadTotalWatts":                    p.LoadTotalWatts,
		"percentageLoad":                    p.PercentageLoad,
		"loadL1Va":                          p.LoadL1Va,
		"loadL2Va":                          p.LoadL2Va,
		"loadL3Va":                          p.LoadL3Va,
		"loadTotalVa":                       p.LoadTotalVa,
		"loadL1Var":                         p.LoadL1Var,
		"loadL2Var":                         p.LoadL2Var,
		"loadL3Var":                         p.LoadL3Var,
		"loadTotalVar":                      p.LoadTotalVar,
		"generatorCumulativeEnergy":         p.GeneratorCumulativeEnergy,
		"generatorCumulativeApparentEnergy": p.GeneratorCumulativeApparentEnergy,
		"generatorCumulativeReactiveEnergy": p.GeneratorCumulativeReactiveEnergy,
		"mainsCumulativeEnergy":             p.MainsCumulativeEnergy,
		"mainsCumulativeApparentEnergy":     p.MainsCumulativeApparentEnergy,
		"mainsCumulativeReactiveEnergy":     p.MainsCumulativeReactiveEnergy,
		"oilPressure":                       p.OilPressure,
		"coolantTemperature":                p.CoolantTemperature,
		"fuelLevel":                         p.FuelLevel,
		"fuelLevelInLit":                    p.FuelLevelInLit,
		"chargeAlternatorVoltage":           p.ChargeAlternatorVoltage,
		"batteryVoltage":                    p.BatteryVoltage,
		"engineSpeed":                       p.EngineSpeed,
		"noOfStarts":                        p.NoOfStarts,
		"noOfTrips":                         p.NoOfTrips,

		"engRunHrs":   p.EngRunHrs,
		"engRunMin":   p.EngRunMin,
		"mainsRunHrs": p.MainsRunHrs,
		"mainsRunMin": p.MainsRunMin,
	}
}

// MqttPayload returns a comma-separated string of all the values
func (p PowerData) MqttPayload() string {
	fields := []string{
		fmt.Sprintf("%d", p.ProtocolRevision),
		fmt.Sprintf("%d", p.GeneratorL1NVoltage),
		fmt.Sprintf("%d", p.GeneratorL2NVoltage),
		fmt.Sprintf("%d", p.GeneratorL3NVoltage),
		fmt.Sprintf("%d", p.GeneratorL1L2Voltage),
		fmt.Sprintf("%d", p.GeneratorL2L3Voltage),
		fmt.Sprintf("%d", p.GeneratorL3L1Voltage),
		fmt.Sprintf("%d", p.GeneratorRFrequency),
		fmt.Sprintf("%d", p.GeneratorYFrequency),
		fmt.Sprintf("%d", p.GeneratorBFrequency),
		fmt.Sprintf("%d", p.GeneratorPowerFactorL1),
		fmt.Sprintf("%d", p.GeneratorPowerFactorL2),
		fmt.Sprintf("%d", p.GeneratorPowerFactorL3),
		fmt.Sprintf("%d", p.GeneratorAveragePowerFactor),
		fmt.Sprintf("%d", p.MainsL1NVoltage),
		fmt.Sprintf("%d", p.MainsL2NVoltage),
		fmt.Sprintf("%d", p.MainsL3NVoltage),
		fmt.Sprintf("%d", p.MainsL1L2Voltage),
		fmt.Sprintf("%d", p.MainsL2L3Voltage),
		fmt.Sprintf("%d", p.MainsL3L1Voltage),
		fmt.Sprintf("%d", p.MainsRFrequency),
		fmt.Sprintf("%d", p.MainsYFrequency),
		fmt.Sprintf("%d", p.MainsBFrequency),
		fmt.Sprintf("%d", p.LoadL1Current),
		fmt.Sprintf("%d", p.LoadL2Current),
		fmt.Sprintf("%d", p.LoadL3Current),
		fmt.Sprintf("%d", p.LoadL1Watts),
		fmt.Sprintf("%d", p.LoadL2Watts),
		fmt.Sprintf("%d", p.LoadL3Watts),
		fmt.Sprintf("%d", p.LoadTotalWatts),
		fmt.Sprintf("%d", p.PercentageLoad),
		fmt.Sprintf("%d", p.LoadL1Va),
		fmt.Sprintf("%d", p.LoadL2Va),
		fmt.Sprintf("%d", p.LoadL3Va),
		fmt.Sprintf("%d", p.LoadTotalVa),
		fmt.Sprintf("%d", p.LoadL1Var),
		fmt.Sprintf("%d", p.LoadL2Var),
		fmt.Sprintf("%d", p.LoadL3Var),
		fmt.Sprintf("%d", p.LoadTotalVar),
		fmt.Sprintf("%d", p.GeneratorCumulativeEnergy),
		fmt.Sprintf("%d", p.GeneratorCumulativeApparentEnergy),
		fmt.Sprintf("%d", p.GeneratorCumulativeReactiveEnergy),
		fmt.Sprintf("%d", p.MainsCumulativeEnergy),
		fmt.Sprintf("%d", p.MainsCumulativeApparentEnergy),
		fmt.Sprintf("%d", p.MainsCumulativeReactiveEnergy),
		fmt.Sprintf("%d", p.OilPressure),
		fmt.Sprintf("%d", p.CoolantTemperature),
		fmt.Sprintf("%d", p.FuelLevel),
		fmt.Sprintf("%d", p.FuelLevelInLit),
		fmt.Sprintf("%d", p.ChargeAlternatorVoltage),
		fmt.Sprintf("%d", p.BatteryVoltage),
		fmt.Sprintf("%d", p.EngineSpeed),
		fmt.Sprintf("%d", p.NoOfStarts),
		fmt.Sprintf("%d", p.NoOfTrips),
		fmt.Sprintf("%d", p.EngRunHrs),
		fmt.Sprintf("%d", p.EngRunMin),
		fmt.Sprintf("%d", p.MainsRunHrs),
		fmt.Sprintf("%d", p.MainsRunMin),
	}

	return strings.Join(fields, ",")
}
