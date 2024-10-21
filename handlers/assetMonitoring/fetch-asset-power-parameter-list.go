package assetMonitoringHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

type Attribute struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Unit  string `json:"unit"`
}

func FetchAssetPowerParameterList(c *fiber.Ctx) error {
	attributes := []Attribute{
		{Label: "Generator L1-N Voltage", Key: "generator_l1n_voltage", Unit: "Volts"},
		{Label: "Generator L2-N Voltage", Key: "generator_l2n_voltage", Unit: "Volts"},
		{Label: "Generator L3-N Voltage", Key: "generator_l3n_voltage", Unit: "Volts"},
		{Label: "Generator L1-L2 Voltage", Key: "generator_l1l2_voltage", Unit: "Volts"},
		{Label: "Generator L2-L3 Voltage", Key: "generator_l2l3_voltage", Unit: "Volts"},
		{Label: "Generator L3-L1 Voltage", Key: "generator_l3l1_voltage", Unit: "Volts"},
		{Label: "Generator R Frequency", Key: "generator_r_frequency", Unit: "Hz"},
		{Label: "Generator Y Frequency", Key: "generator_y_frequency", Unit: "Hz"},
		{Label: "Generator B Frequency", Key: "generator_b_frequency", Unit: "Hz"},
		{Label: "Generator Power Factor L1", Key: "generator_power_factor_l1"},
		{Label: "Generator Power Factor L2", Key: "generator_power_factor_l2"},
		{Label: "Generator Power Factor L3", Key: "generator_power_factor_l3"},
		{Label: "Generator Average Power Factor", Key: "generator_average_power_factor"},
		{Label: "Mains L1-N Voltage", Key: "mains_l1n_voltage", Unit: "Volts"},
		{Label: "Mains L2-N Voltage", Key: "mains_l2n_voltage", Unit: "Volts"},
		{Label: "Mains L3-N Voltage", Key: "mains_l3n_voltage", Unit: "Volts"},
		{Label: "Mains L1-L2 Voltage", Key: "mains_l1l2_voltage", Unit: "Volts"},
		{Label: "Mains L2-L3 Voltage", Key: "mains_l2l3_voltage", Unit: "Volts"},
		{Label: "Mains L3-L1 Voltage", Key: "mains_l3l1_voltage", Unit: "Volts"},
		{Label: "Mains R Frequency", Key: "mains_r_frequency", Unit: "Hz"},
		{Label: "Mains Y Frequency", Key: "mains_y_frequency", Unit: "Hz"},
		{Label: "Mains B Frequency", Key: "mains_b_frequency", Unit: "Hz"},
		{Label: "Load L1 Current", Key: "load_l1_current", Unit: "Amps"},
		{Label: "Load L2 Current", Key: "load_l2_current", Unit: "Amps"},
		{Label: "Load L3 Current", Key: "load_l3_current", Unit: "Amps"},
		{Label: "Load L1 Watts", Key: "load_l1_watts", Unit: "KW"},
		{Label: "Load L2 Watts", Key: "load_l2_watts", Unit: "KW"},
		{Label: "Load L3 Watts", Key: "load_l3_watts", Unit: "KW"},
		{Label: "Load Total Watts", Key: "load_total_watts", Unit: "KW"},
		{Label: "Percentage Load", Key: "percentage_load", Unit: "%"},
		{Label: "Load L1 VA", Key: "load_l1_va", Unit: "VA"},
		{Label: "Load L2 VA", Key: "load_l2_va", Unit: "VA"},
		{Label: "Load L3 VA", Key: "load_l3_va", Unit: "VA"},
		{Label: "Load Total VA", Key: "load_total_va", Unit: "VA"},
		{Label: "Load L1 VAR", Key: "load_l1_var", Unit: "VAR"},
		{Label: "Load L2 VAR", Key: "load_l2_var", Unit: "VAR"},
		{Label: "Load L3 VAR", Key: "load_l3_var", Unit: "VAR"},
		{Label: "Load Total VAR", Key: "load_total_var", Unit: "VAR"},
		{Label: "Generator Cumulative Energy", Key: "generator_cumulative_energy", Unit: "KWH"},
		{Label: "Generator Cumulative Apparent Energy", Key: "generator_cumulative_apparent_energy", Unit: "KWAH"},
		{Label: "Generator Cumulative Reactive Energy", Key: "generator_cumulative_reactive_energy", Unit: "KWARH"},
		{Label: "Mains Cumulative Energy", Key: "mains_cumulative_energy", Unit: "KWH"},
		{Label: "Mains Cumulative Apparent Energy", Key: "mains_cumulative_apparent_energy", Unit: "KWAH"},
		{Label: "Mains Cumulative Reactive Energy", Key: "mains_cumulative_reactive_energy", Unit: "KWARH"},
		{Label: "Oil Pressure", Key: "oil_pressure", Unit: "Bar"},
		{Label: "Coolant Temperature", Key: "coolant_temperature", Unit: "°C"},
		{Label: "Fuel Level", Key: "fuel_level", Unit: "%"},
		{Label: "Fuel Level in Liters", Key: "fuel_level_in_lit", Unit: "Litre"},
		{Label: "Charge Alternator Voltage", Key: "charge_alternator_voltage", Unit: "Volts"},
		{Label: "Battery Voltage", Key: "battery_voltage", Unit: "Volts"},
		{Label: "Engine Speed", Key: "engine_speed", Unit: "rpm"},
		{Label: "Number of Starts", Key: "no_of_starts"},
		{Label: "Number of Trips", Key: "no_of_trips"},
		{Label: "Engine Run Hours", Key: "eng_run_hrs", Unit: "Hrs"},
		{Label: "Engine Run Minutes", Key: "eng_run_min", Unit: "mins"},
		{Label: "Mains Run Hours", Key: "mains_run_hrs", Unit: "Hrs"},
		{Label: "Mains Run Minutes", Key: "mains_run_min", Unit: "mins"},
	}

	payload := fiber.Map{
		"powerParameters": attributes,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
