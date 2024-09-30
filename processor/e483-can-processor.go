package processor

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessE483Can(client MQTT.Client, serialNo, message string) {
	device := database.FindDeviceBySerialNo(serialNo)
	if device.IsIdNull() {
		log.Error().Str("ProcessE483Can: Device not found with serial no", serialNo).Send()
		return
	}

	canDataStr := strings.Split(message, ",")
	canData := []float64{}

	for _, canDataItem := range canDataStr {
		dataFloat64, err := strconv.ParseFloat(canDataItem, 64)
		if err != nil {
			log.Error().AnErr("ProcessE483Can: Parse ", err).Send()
			return
		}
		canData = append(canData, float64(dataFloat64))
	}

	if len(canData) != 58 {
		log.Error().Str("ProcessE483Can: Wrong can data", serialNo).Send()
		return
	}

	newCanData := models.E483CanData{
		DeviceId: device.Id,
	}
	newCanData.CoolantTemperature = uint8(canData[0])
	newCanData.FuelTemperature = uint8(canData[1])
	newCanData.StarterMode = uint8(canData[2])
	newCanData.RPM = float32(canData[3])
	newCanData.EngineDemandPercentTorque = int8(canData[4])
	newCanData.ActualPercentTorqueHighRes = float32(canData[5])
	newCanData.DriverDemandPercentTorque = int8(canData[6])
	newCanData.ActualEnginePercentTorque = int8(canData[7])
	newCanData.TorqueMode = uint8(canData[8])
	newCanData.HoursByEMS = float64(canData[9])
	newCanData.WaterInFuel = uint8(canData[10])
	newCanData.FuelEconomy = float32(canData[11])
	newCanData.FuelConsumptionTotalHighRes = float64(canData[12])
	newCanData.BatteryPotential = float32(canData[13])
	newCanData.UreaPercentageInWater = float32(canData[14])
	newCanData.AdblueTankLevel = float32(canData[15])
	newCanData.EngineOilPressure = uint16(canData[16])

	err := database.Database.Create(&newCanData).Error
	if err != nil {
		log.Error().AnErr("ProcessE483Can: create canData", err).Send()
		return
	}

	client.Publish("iisc/web/"+device.SerialNo+"/e483can", 0, false, newCanData.MqttPayload())
}
