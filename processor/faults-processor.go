package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func ProcessFaults(client MQTT.Client, topic string, message string) {
	deviceId, err := processFaultsTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	var deviceFault models.DeviceFault

	rawStringArr := strings.Split(message, ",")
	if len(rawStringArr) != 33 {
		return
	}

	deviceFault.EhmDeviceId = &ehmDevice.Id
	deviceFault.LowOilPressure = rawStringArr[0] == "1"
	deviceFault.HighCoolantTemperature = rawStringArr[1] == "1"
	deviceFault.LowFuelLevel = rawStringArr[2] == "1"
	deviceFault.WaterLevelSwitch = rawStringArr[3] == "1"
	deviceFault.UnderSpeed = rawStringArr[4] == "1"
	deviceFault.OverSpeed = rawStringArr[5] == "1"
	deviceFault.FailToStart = rawStringArr[6] == "1"
	deviceFault.FailToStop = rawStringArr[7] == "1"
	deviceFault.ReversePowerDetected = rawStringArr[8] == "1"
	deviceFault.LowLoadAlarm = rawStringArr[9] == "1"
	deviceFault.GeneratorLowFrequency = rawStringArr[10] == "1"
	deviceFault.GeneratorHighFrequency = rawStringArr[11] == "1"
	deviceFault.GeneratorHighCurrent = rawStringArr[12] == "1"
	deviceFault.GeneratorOverload = rawStringArr[13] == "1"
	deviceFault.UnbalancedLoad = rawStringArr[14] == "1"
	deviceFault.EmergencyStrip = rawStringArr[15] == "1"
	deviceFault.ChargeAlternatorFailure = rawStringArr[16] == "1"
	deviceFault.MaintenanceDue = rawStringArr[17] == "1"
	deviceFault.AftActivationTimeout = rawStringArr[18] == "1"
	deviceFault.AshLoadMaintainance = rawStringArr[19] == "1"
	deviceFault.BatteryLowVoltage = rawStringArr[20] == "1"
	deviceFault.BatteryHighVoltage = rawStringArr[21] == "1"
	deviceFault.TemperatureCircuitOpen = rawStringArr[22] == "1"
	deviceFault.LopBatteryShort = rawStringArr[23] == "1"
	deviceFault.FuelTheft = rawStringArr[24] == "1"
	deviceFault.MagneticPickUpFault = rawStringArr[25] == "1"
	deviceFault.OilPressureOpenCircuit = rawStringArr[26] == "1"
	deviceFault.L1LowVoltage = rawStringArr[27] == "1"
	deviceFault.L1HighVoltage = rawStringArr[28] == "1"
	deviceFault.L2LowVoltage = rawStringArr[29] == "1"
	deviceFault.L2HighVoltage = rawStringArr[30] == "1"
	deviceFault.L3LowVoltage = rawStringArr[31] == "1"
	deviceFault.L3HighVoltage = rawStringArr[32] == "1"

	err = database.Database.Create(&deviceFault).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/faults/%s", deviceFault.EhmDeviceId)
	dataToSend, err := json.Marshal(deviceFault.Json())
	if err != nil {
		log.Errorln(err.Error())
	} else {
		log.Println("Publishing to topic: ", publishTopic)
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}
func processFaultsTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("fuel percentage topic invalid length")
	}

	return rawStringArr[0], nil
}
