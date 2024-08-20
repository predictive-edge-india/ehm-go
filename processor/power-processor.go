package processor

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessPowerData(client MQTT.Client, serialNo, message string) {
	device := database.FindDeviceBySerialNo(serialNo)
	if device.IsIdNull() {
		log.Error().Str("ProcessPowerData: Device not found with serial no", serialNo).Send()
		return
	}

	powerDataStr := strings.Split(message, ",")
	powerData := []int32{}

	for _, powerDataItem := range powerDataStr {
		dataFloat32, err := strconv.ParseFloat(powerDataItem, 32)
		if err != nil {
			log.Error().AnErr("ProcessPowerData: Parse ", err).Send()
			return
		}
		powerData = append(powerData, int32(dataFloat32))
	}

	if len(powerData) != 58 {
		log.Error().Str("ProcessPowerData: Wrong powerData", serialNo).Send()
		return
	}

	newPowerData := models.PowerData{
		DeviceId: device.Id,
	}

	newPowerData.ProtocolRevision = uint16(powerData[0])
	newPowerData.GeneratorL1NVoltage = uint16(powerData[1])
	newPowerData.GeneratorL2NVoltage = uint16(powerData[2])
	newPowerData.GeneratorL3NVoltage = uint16(powerData[3])
	newPowerData.GeneratorL1L2Voltage = uint16(powerData[4])
	newPowerData.GeneratorL2L3Voltage = uint16(powerData[5])
	newPowerData.GeneratorL3L1Voltage = uint16(powerData[6])

	newPowerData.GeneratorRFrequency = uint16(powerData[7])
	newPowerData.GeneratorYFrequency = uint16(powerData[8])
	newPowerData.GeneratorBFrequency = uint16(powerData[9])

	newPowerData.GeneratorPowerFactorL1 = int16(powerData[10])
	newPowerData.GeneratorPowerFactorL2 = int16(powerData[11])
	newPowerData.GeneratorPowerFactorL3 = int16(powerData[12])
	newPowerData.GeneratorAveragePowerFactor = int16(powerData[13])

	newPowerData.MainsL1NVoltage = uint16(powerData[14])
	newPowerData.MainsL2NVoltage = uint16(powerData[15])
	newPowerData.MainsL3NVoltage = uint16(powerData[16])
	newPowerData.MainsL1L2Voltage = uint16(powerData[17])
	newPowerData.MainsL2L3Voltage = uint16(powerData[18])
	newPowerData.MainsL3L1Voltage = uint16(powerData[19])

	newPowerData.MainsRFrequency = uint16(powerData[20])
	newPowerData.MainsYFrequency = uint16(powerData[21])
	newPowerData.MainsBFrequency = uint16(powerData[22])

	newPowerData.LoadL1Current = uint16(powerData[23])
	newPowerData.LoadL2Current = uint16(powerData[24])
	newPowerData.LoadL3Current = uint16(powerData[25])

	newPowerData.LoadL1Watts = int16(powerData[26])
	newPowerData.LoadL2Watts = int16(powerData[27])
	newPowerData.LoadL3Watts = int16(powerData[28])
	newPowerData.LoadTotalWatts = int16(powerData[29])
	newPowerData.PercentageLoad = int16(powerData[30])

	newPowerData.LoadL1Va = uint16(powerData[31])
	newPowerData.LoadL2Va = uint16(powerData[32])
	newPowerData.LoadL3Va = uint16(powerData[33])
	newPowerData.LoadTotalVa = uint16(powerData[34])

	newPowerData.LoadL1Var = uint16(powerData[35])
	newPowerData.LoadL2Var = uint16(powerData[36])
	newPowerData.LoadL3Var = uint16(powerData[37])
	newPowerData.LoadTotalVar = uint16(powerData[38])

	newPowerData.GeneratorCumulativeEnergy = uint32(powerData[39])
	newPowerData.GeneratorCumulativeApparentEnergy = uint32(powerData[40])
	newPowerData.GeneratorCumulativeReactiveEnergy = uint32(powerData[41])
	newPowerData.MainsCumulativeEnergy = uint32(powerData[42])
	newPowerData.MainsCumulativeApparentEnergy = uint32(powerData[43])
	newPowerData.MainsCumulativeReactiveEnergy = uint32(powerData[44])

	newPowerData.OilPressure = uint16(powerData[45])
	newPowerData.CoolantTemperature = int16(powerData[46])
	newPowerData.FuelLevel = uint16(powerData[47])

	newPowerData.FuelLevelInLit = uint16(powerData[48])
	newPowerData.ChargeAlternatorVoltage = uint16(powerData[49])
	newPowerData.BatteryVoltage = uint16(powerData[50])
	newPowerData.EngineSpeed = uint16(powerData[51])
	newPowerData.NoOfStarts = uint16(powerData[52])
	newPowerData.NoOfTrips = uint16(powerData[53])
	newPowerData.EngRunHrs = uint32(powerData[54])
	newPowerData.EngRunMin = uint16(powerData[55])
	newPowerData.MainsRunHrs = uint16(powerData[56])
	newPowerData.MainsRunMin = uint16(powerData[57])

	if newPowerData.ProtocolRevision < 1 {
		log.Error().Msg("ProcessPowerData: Invalid data")
		return
	}

	err := database.Database.Create(&newPowerData).Error
	if err != nil {
		log.Error().AnErr("ProcessPowerData: create newPowerData", err).Send()
		return
	}

	client.Publish("iisc/web/"+device.SerialNo+"/power", 0, false, newPowerData.MqttPayload())
}
