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

func ProcessTemperature(client MQTT.Client, topic string, message string) {
	deviceId, err := processTemperatureTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	var temperatureParam models.TemperatureParam

	rawStringArr := strings.Split(message, ",")
	if len(rawStringArr) != 3 {
		return
	}

	var temperatures []float32

	for i := 0; i < len(rawStringArr); i++ {
		temperature := helpers.ParseFloat32(rawStringArr[i])
		temperatures = append(temperatures, temperature)
	}

	temperatureParam.EhmDeviceId = &ehmDevice.Id
	temperatureParam.Temperatures = temperatures

	err = database.Database.Create(&temperatureParam).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/%s/temperature", temperatureParam.EhmDeviceId)
	dataToSend, err := json.Marshal(temperatureParam.Json())
	if err != nil {
		log.Errorln(err.Error())
	} else {
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}

func processTemperatureTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("temperature topic invalid length")
	}

	return rawStringArr[0], nil
}
