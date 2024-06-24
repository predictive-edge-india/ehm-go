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

	var powerParam models.PowerParam

	rawStringArr := strings.Split(message, ",")
	if len(rawStringArr) != 22 {
		return
	}

	powerParam.EhmDeviceId = &ehmDevice.Id
	for i := 0; i < len(rawStringArr); i++ {
		switch i {
		case 0:
			powerParam.L1NVoltage = helpers.StringToUint16(rawStringArr[i])
		case 1:
			powerParam.L2NVoltage = helpers.StringToUint16(rawStringArr[i])
		case 2:
			powerParam.L3NVoltage = helpers.StringToUint16(rawStringArr[i])
		case 3:
			powerParam.L1L2Voltage = helpers.StringToUint16(rawStringArr[i])
		case 4:
			powerParam.L2L3Voltage = helpers.StringToUint16(rawStringArr[i])
		case 5:
			powerParam.L2L1Voltage = helpers.StringToUint16(rawStringArr[i])
		case 6:
			powerParam.L1Current = helpers.StringToUint16(rawStringArr[i])
		case 7:
			powerParam.L2Current = helpers.StringToUint16(rawStringArr[i])
		case 8:
			powerParam.L3Current = helpers.StringToUint16(rawStringArr[i])
		case 9:
			powerParam.TotalWatts = helpers.StringToUint16(rawStringArr[i])
		case 10:
			powerParam.LoadPercentage = helpers.StringToUint16(rawStringArr[i])
		case 11:
			powerParam.L1VA = helpers.StringToUint16(rawStringArr[i])
		case 12:
			powerParam.L2VA = helpers.StringToUint16(rawStringArr[i])
		case 13:
			powerParam.L3VA = helpers.StringToUint16(rawStringArr[i])
		case 14:
			powerParam.TotalVA = helpers.StringToUint16(rawStringArr[i])
		case 15:
			powerParam.RFrequency = helpers.StringToUint8(rawStringArr[i])
		case 16:
			powerParam.YFrequency = helpers.StringToUint8(rawStringArr[i])
		case 17:
			powerParam.BFrequency = helpers.StringToUint8(rawStringArr[i])
		case 18:
			powerParam.PfL1 = helpers.ParseFloat32(rawStringArr[i])
		case 19:
			powerParam.PfL2 = helpers.ParseFloat32(rawStringArr[i])
		case 20:
			powerParam.PfL3 = helpers.ParseFloat32(rawStringArr[i])
		case 21:
			powerParam.PfAvg = helpers.ParseFloat32(rawStringArr[i])
		}
	}

	err = database.Database.Create(&powerParam).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/power/%s", powerParam.EhmDeviceId)
	dataToSend, err := json.Marshal(powerParam.Json())
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

func processPowerParamTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("power param topic invalid length")
	}

	return rawStringArr[0], nil
}
