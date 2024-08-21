package assetMonitoringHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/predictive-edge-india/ehm-go/helpers"
)

type Attribute struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func FetchAssetPowerParameterList(c *fiber.Ctx) error {
	attributes := []Attribute{
		{Label: "Generator L1-N Voltage", Key: "generator_l1n_voltage"},
		{Label: "Generator L2-N Voltage", Key: "generator_l2n_voltage"},
		{Label: "Generator L3-N Voltage", Key: "generator_l3n_voltage"},
		{Label: "Generator L1-L2 Voltage", Key: "generator_l1l2_voltage"},
		{Label: "Generator L2-L3 Voltage", Key: "generator_l2l3_voltage"},
		{Label: "Generator L3-L1 Voltage", Key: "generator_l3l1_voltage"},
		{Label: "Generator R Frequency", Key: "generator_r_frequency"},
		{Label: "Generator Y Frequency", Key: "generator_y_frequency"},
		{Label: "Generator B Frequency", Key: "generator_b_frequency"},
		{Label: "Generator Power Factor L1", Key: "generator_power_factor_l1"},
		{Label: "Generator Power Factor L2", Key: "generator_power_factor_l2"},
		{Label: "Generator Power Factor L3", Key: "generator_power_factor_l3"},
		{Label: "Generator Average Power Factor", Key: "generator_average_power_factor"},
		{Label: "Mains L1-N Voltage", Key: "mains_l1n_voltage"},
		{Label: "Mains L2-N Voltage", Key: "mains_l2n_voltage"},
		{Label: "Mains L3-N Voltage", Key: "mains_l3n_voltage"},
		{Label: "Mains L1-L2 Voltage", Key: "mains_l1l2_voltage"},
		{Label: "Mains L2-L3 Voltage", Key: "mains_l2l3_voltage"},
		{Label: "Mains L3-L1 Voltage", Key: "mains_l3l1_voltage"},
		{Label: "Mains R Frequency", Key: "mains_r_frequency"},
		{Label: "Mains Y Frequency", Key: "mains_y_frequency"},
		{Label: "Mains B Frequency", Key: "mains_b_frequency"},
		{Label: "Load L1 Current", Key: "load_l1_current"},
		{Label: "Load L2 Current", Key: "load_l2_current"},
		{Label: "Load L3 Current", Key: "load_l3_current"},
		{Label: "Load L1 Watts", Key: "load_l1_watts"},
		{Label: "Load L2 Watts", Key: "load_l2_watts"},
		{Label: "Load L3 Watts", Key: "load_l3_watts"},
		{Label: "Load Total Watts", Key: "load_total_watts"},
		{Label: "Percentage Load", Key: "percentage_load"},
		{Label: "Load L1 VA", Key: "load_l1_va"},
		{Label: "Load L2 VA", Key: "load_l2_va"},
		{Label: "Load L3 VA", Key: "load_l3_va"},
		{Label: "Load Total VA", Key: "load_total_va"},
		{Label: "Load L1 VAR", Key: "load_l1_var"},
		{Label: "Load L2 VAR", Key: "load_l2_var"},
		{Label: "Load L3 VAR", Key: "load_l3_var"},
		{Label: "Load Total VAR", Key: "load_total_var"},
		{Label: "Generator Cumulative Energy", Key: "generator_cumulative_energy"},
		{Label: "Generator Cumulative Apparent Energy", Key: "generator_cumulative_apparent_energy"},
		{Label: "Generator Cumulative Reactive Energy", Key: "generator_cumulative_reactive_energy"},
		{Label: "Mains Cumulative Energy", Key: "mains_cumulative_energy"},
		{Label: "Mains Cumulative Apparent Energy", Key: "mains_cumulative_apparent_energy"},
		{Label: "Mains Cumulative Reactive Energy", Key: "mains_cumulative_reactive_energy"},
		{Label: "Oil Pressure", Key: "oil_pressure"},
		{Label: "Coolant Temperature", Key: "coolant_temperature"},
		{Label: "Fuel Level", Key: "fuel_level"},
		{Label: "Fuel Level in Liters", Key: "fuel_level_in_lit"},
		{Label: "Charge Alternator Voltage", Key: "charge_alternator_voltage"},
		{Label: "Battery Voltage", Key: "battery_voltage"},
		{Label: "Engine Speed", Key: "engine_speed"},
		{Label: "Number of Starts", Key: "no_of_starts"},
		{Label: "Number of Trips", Key: "no_of_trips"},
		{Label: "Engine Run Hours", Key: "eng_run_hrs"},
		{Label: "Engine Run Minutes", Key: "eng_run_min"},
		{Label: "Mains Run Hours", Key: "mains_run_hrs"},
		{Label: "Mains Run Minutes", Key: "mains_run_min"},
	}

	payload := fiber.Map{
		"powerParameters": attributes,
	}

	return c.JSON(helpers.BuildResponse(payload))
}
