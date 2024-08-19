package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlarmStatusFlag struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`

	LowOilPressure          uint8 `gorm:"column:low_oil_pressure;type:tinyint(1) unsigned;not null"`
	HighCoolantTemperature  uint8 `gorm:"column:high_coolant_temperature;type:tinyint(1) unsigned;not null"`
	LowFuelLevel            uint8 `gorm:"column:low_fuel_level;type:tinyint(1) unsigned;not null"`
	WaterLevelSwitch        uint8 `gorm:"column:water_level_switch;type:tinyint(1) unsigned;not null"`
	UnderSpeed              uint8 `gorm:"column:under_speed;type:tinyint(1) unsigned;not null"`
	OverSpeed               uint8 `gorm:"column:over_speed;type:tinyint(1) unsigned;not null"`
	FailToStart             uint8 `gorm:"column:fail_to_start;type:tinyint(1) unsigned;not null"`
	FailToStop              uint8 `gorm:"column:fail_to_stop;type:tinyint(1) unsigned;not null"`
	ReversePowerDetected    uint8 `gorm:"column:reverse_power_detected;type:tinyint(1) unsigned;not null"`
	LowLoadAlarm            uint8 `gorm:"column:low_load_alarm;type:tinyint(1) unsigned;not null"`
	GeneratorLowFrequency   uint8 `gorm:"column:generator_low_frequency;type:tinyint(1) unsigned;not null"`
	GeneratorHighFrequency  uint8 `gorm:"column:generator_high_frequency;type:tinyint(1) unsigned;not null"`
	GeneratorHighCurrent    uint8 `gorm:"column:generator_high_current;type:tinyint(1) unsigned;not null"`
	GeneratorOverload       uint8 `gorm:"column:generator_overload;type:tinyint(1) unsigned;not null"`
	UnbalancedLoad          uint8 `gorm:"column:unbalanced_load;type:tinyint(1) unsigned;not null"`
	EmergencyStop           uint8 `gorm:"column:emergency_stop;type:tinyint(1) unsigned;not null"`
	ChargeAlternatorFailure uint8 `gorm:"column:charge_alternator_failure;type:tinyint(1) unsigned;not null"`
	MaintenanceDue          uint8 `gorm:"column:maintenance_due;type:tinyint(1) unsigned;not null"`
	AftActivationTimeout    uint8 `gorm:"column:aft_activation_timeout;type:tinyint(1) unsigned;not null"`
	AshLoadMaintenance      uint8 `gorm:"column:ash_load_maintenance;type:tinyint(1) unsigned;not null"`
	BatteryLowVoltage       uint8 `gorm:"column:battery_low_voltage;type:tinyint(1) unsigned;not null"`
	BatteryHighVoltage      uint8 `gorm:"column:battery_high_voltage;type:tinyint(1) unsigned;not null"`
	TemperatureCircuitOpen  uint8 `gorm:"column:temperature_circuit_open;type:tinyint(1) unsigned;not null"`
	LopPin23ShortToBattery  uint8 `gorm:"column:lop_pin23_short_to_battery;type:tinyint(1) unsigned;not null"`
	FuelTheft               uint8 `gorm:"column:fuel_theft;type:tinyint(1) unsigned;not null"`
	MagneticPickUpFault     uint8 `gorm:"column:magnetic_pick_up_fault;type:tinyint(1) unsigned;not null"`
	OilPressureOpenCircuit  uint8 `gorm:"column:oil_pressure_open_circuit;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputI         uint8 `gorm:"column:auxiliary_input_i;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputA         uint8 `gorm:"column:auxiliary_input_a;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputB         uint8 `gorm:"column:auxiliary_input_b;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputC         uint8 `gorm:"column:auxiliary_input_c;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputD         uint8 `gorm:"column:auxiliary_input_d;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputE         uint8 `gorm:"column:auxiliary_input_e;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputF         uint8 `gorm:"column:auxiliary_input_f;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputG         uint8 `gorm:"column:auxiliary_input_g;type:tinyint(1) unsigned;not null"`
	AuxiliaryInputH         uint8 `gorm:"column:auxiliary_input_h;type:tinyint(1) unsigned;not null"`
	GenL1PhaseLowVolt       uint8 `gorm:"column:gen_l1_phase_low_volt;type:tinyint(1) unsigned;not null"`
	GenL1PhaseHighVolt      uint8 `gorm:"column:gen_l1_phase_high_volt;type:tinyint(1) unsigned;not null"`
	GenL2PhaseLowVolt       uint8 `gorm:"column:gen_l2_phase_low_volt;type:tinyint(1) unsigned;not null"`
	GenL2PhaseHighVolt      uint8 `gorm:"column:gen_l2_phase_high_volt;type:tinyint(1) unsigned;not null"`
	GenL3PhaseLowVolt       uint8 `gorm:"column:gen_l3_phase_low_volt;type:tinyint(1) unsigned;not null"`
	GenL3PhaseHighVolt      uint8 `gorm:"column:gen_l3_phase_high_volt;type:tinyint(1) unsigned;not null"`
	DgPhaseRotation         uint8 `gorm:"column:dg_phase_rotation;type:tinyint(1) unsigned;not null"`
	MainsPhaseRotation      uint8 `gorm:"column:mains_phase_rotation;type:tinyint(1) unsigned;not null"`
	FuelLevelOpenCircuit    uint8 `gorm:"column:fuel_level_open_circuit;type:tinyint(1) unsigned;not null"`
	VBeltBroken             uint8 `gorm:"column:v_belt_broken;type:tinyint(1) unsigned;not null"`
	Reserved                uint8 `gorm:"column:reserved;type:tinyint(1) unsigned;not null"`
	HighOilPressureDetected uint8 `gorm:"column:high_oil_pressure_detected;type:tinyint(1) unsigned;not null"`
}

func (AlarmStatusFlag) TableName() string {
	return "alarm_status_flags"
}
