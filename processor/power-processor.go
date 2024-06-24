package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/helpers"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func ProcessPowerParam(client MQTT.Client, topic string, message string) {
	deviceId, err := processPowerParamTopic(topic)
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
	if len(rawStringArr) != 22 {
		return
	}

	powerParam := models.PowerParam{
		EhmDeviceId:    &ehmDevice.Id,
		L1NVoltage:     helpers.StringToUint16(rawStringArr[0]),
		L2NVoltage:     helpers.StringToUint16(rawStringArr[1]),
		L3NVoltage:     helpers.StringToUint16(rawStringArr[2]),
		L1L2Voltage:    helpers.StringToUint16(rawStringArr[3]),
		L2L3Voltage:    helpers.StringToUint16(rawStringArr[4]),
		L2L1Voltage:    helpers.StringToUint16(rawStringArr[5]),
		RFrequency:     helpers.StringToUint8(rawStringArr[6]),
		YFrequency:     helpers.StringToUint8(rawStringArr[7]),
		BFrequency:     helpers.StringToUint8(rawStringArr[8]),
		PfL1:           helpers.ParseFloat32(rawStringArr[9]),
		PfL2:           helpers.ParseFloat32(rawStringArr[10]),
		PfL3:           helpers.ParseFloat32(rawStringArr[11]),
		PfAvg:          helpers.ParseFloat32(rawStringArr[12]),
		L1Current:      helpers.StringToUint16(rawStringArr[13]),
		L2Current:      helpers.StringToUint16(rawStringArr[14]),
		L3Current:      helpers.StringToUint16(rawStringArr[15]),
		TotalWatts:     helpers.StringToUint16(rawStringArr[16]),
		LoadPercentage: helpers.StringToUint16(rawStringArr[17]),
		L1VA:           helpers.StringToUint16(rawStringArr[18]),
		L2VA:           helpers.StringToUint16(rawStringArr[19]),
		L3VA:           helpers.StringToUint16(rawStringArr[20]),
		TotalVA:        helpers.StringToUint16(rawStringArr[21]),
	}

	err = database.Database.Create(&powerParam).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/%s/power", powerParam.EhmDeviceId)
	dataToSend, err := json.Marshal(powerParam.Json())
	if err != nil {
		log.Errorln(err.Error())
	} else {
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}

func processPowerParamTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("power param topic invalid length")
	}

	return rawStringArr[0], nil
}
