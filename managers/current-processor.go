package managers

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

var gRms float64

func ProcessCurrentMessage(client MQTT.Client, topic string, message string) {
	newParam := new(models.CurrentParameter)

	deviceId, paramType, dataType, err := processCurrentTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	newParam.ParamType = paramType
	newParam.EhmDeviceId = &ehmDevice.Id

	if dataType == models.InputDataType.RMS {
		rmsValue, err := parseRms(message)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		newParam.RMS = helpers.ToFixed(rmsValue, 2)
	} else if dataType == models.InputDataType.FFT {
		fftValues := ParseFft(message)
		newParam.FFT = fftValues
	}

	newParam.PacketType = dataType

	if newParam.PacketType == models.InputDataType.RMS {
		gRms = newParam.RMS
	}
	if newParam.PacketType == models.InputDataType.FFT && gRms > 0 {
		newParam.RMS = gRms
		newParam = database.CreateCurrentParameter(*newParam)

		publishTopic := fmt.Sprintf("iisc/web/graph/%s/%s", newParam.EhmDeviceId, newParam.ParamType)

		dataToSend, err := json.Marshal(newParam.Json())
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
}

func processCurrentTopic(topic string) (string, string, int32, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 3 {
		return "", "", -1, errors.New("current topic invalid length")
	}

	if rawStringArr[2] == "rms" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.RMS, nil
	}
	if rawStringArr[2] == "fft" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.FFT, nil
	}

	return "", "", -1, errors.New("current topic invalid type")
}
