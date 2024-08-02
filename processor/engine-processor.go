package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	log "github.com/sirupsen/logrus"
)

func ProcessEngineParam(client MQTT.Client, topic string, message string) {
	deviceId, err := processEngineParamTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	rawStringArr := strings.Split(message, ",")
	if len(rawStringArr) != 30 {
		return
	}

	engineParam := models.EngineParam{
		EhmDeviceId:                    &ehmDevice.Id,
		CoolantTemperature:             helpers.StringToUint8(rawStringArr[0]),
		FuelTemperature:                helpers.StringToUint8(rawStringArr[1]),
		StarterMode:                    helpers.StringToUint8(rawStringArr[2]),
		Rpm:                            helpers.ParseFloat32(rawStringArr[3]),
		EngineDemandPercentTorque:      helpers.StringToInt8(rawStringArr[4]),
		ActualPercentTorqueHighRes:     helpers.ParseFloat32(rawStringArr[5]),
		DriverDemandPercentTorque:      helpers.StringToInt8(rawStringArr[6]),
		ActualEnginePercentTorque:      helpers.StringToInt8(rawStringArr[7]),
		TorqueMode:                     helpers.StringToUint8(rawStringArr[8]),
		HoursByEms:                     helpers.ParseFloat32(rawStringArr[9]),
		Second:                         helpers.ParseFloat32(rawStringArr[10]),
		Minute:                         helpers.StringToUint8(rawStringArr[11]),
		Hour:                           helpers.StringToUint8(rawStringArr[12]),
		Month:                          helpers.StringToUint8(rawStringArr[13]),
		Day:                            helpers.StringToUint8(rawStringArr[14]),
		Year:                           helpers.StringToUint16(rawStringArr[15]),
		RequestedSpeed:                 helpers.ParseFloat32(rawStringArr[16]),
		RequestedTorque:                helpers.StringToInt8(rawStringArr[17]),
		RequestedTorqueHighRes:         helpers.ParseFloat32(rawStringArr[18]),
		WaterInFuel:                    helpers.StringToUint8(rawStringArr[19]),
		InducementLevel:                helpers.StringToUint8(rawStringArr[20]),
		FuelEconomy:                    helpers.ParseFloat32(rawStringArr[21]),
		FuelConsumptionTotalHighRes:    helpers.ParseFloat32(rawStringArr[22]),
		BatteryPotential:               helpers.ParseFloat32(rawStringArr[23]),
		MalfunctionLampIndicatorStatus: helpers.StringToUint8(rawStringArr[24]),
		StopLampStatus:                 helpers.StringToUint8(rawStringArr[25]),
		WarningLampStatus:              helpers.StringToUint8(rawStringArr[26]),
		UreaPercentageInWater:          helpers.ParseFloat32(rawStringArr[27]),
		AdblueTankLevel:                helpers.ParseFloat32(rawStringArr[28]),
		ParkBrakeStatus:                helpers.StringToUint8(rawStringArr[29]),
	}

	err = database.Database.Create(&engineParam).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/%s/engine", engineParam.EhmDeviceId)
	dataToSend, err := json.Marshal(engineParam.Json())
	if err != nil {
		log.Errorln(err.Error())
	} else {
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}

func processEngineParamTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("engine param topic invalid length")
	}

	return rawStringArr[0], nil
}
