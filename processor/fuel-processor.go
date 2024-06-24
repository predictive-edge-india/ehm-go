package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func ProcessFuelPercentage(client MQTT.Client, topic string, message string) {
	deviceId, err := processFuelPercentageTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	var fuelPercentage models.FuelPercentage
	percent, err := strconv.ParseInt(message, 10, 32)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	fuelPercentage.EhmDeviceId = &ehmDevice.Id
	fuelPercentage.Percentage = int16(percent)
	err = database.Database.Create(&fuelPercentage).Error
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/%s/fuel", fuelPercentage.EhmDeviceId)
	dataToSend, err := json.Marshal(fuelPercentage.Json())
	if err != nil {
		log.Errorln(err.Error())
	} else {
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}

func processFuelPercentageTopic(topic string) (string, error) {
	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 2 {
		return "", errors.New("fuel percentage topic invalid length")
	}

	return rawStringArr[0], nil
}
