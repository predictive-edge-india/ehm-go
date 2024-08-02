package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessFuelPercentage(client MQTT.Client, topic string, message string) {
	deviceId, err := processFuelPercentageTopic(topic)
	if err != nil {
		log.Error().AnErr("ProcessFuelPercentage: processFuelPercentageTopic", err).Send()
		return
	}
	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Error().AnErr("ProcessFuelPercentage: FindOrCreateEhmDevice", err).Send()
		return
	}

	var fuelPercentage models.FuelPercentage
	percent, err := strconv.ParseInt(message, 10, 32)
	if err != nil {
		log.Error().AnErr("ProcessFuelPercentage: ParseInt", err).Send()
		return
	}

	fuelPercentage.EhmDeviceId = &ehmDevice.Id
	fuelPercentage.Percentage = int16(percent)
	if err = database.Database.Create(&fuelPercentage).Error; err != nil {
		log.Error().AnErr("ProcessFuelPercentage: create fuelPercentage", err).Send()
		return
	}

	publishTopic := fmt.Sprintf("iisc/web/%s/fuel", fuelPercentage.EhmDeviceId)
	dataToSend, err := json.Marshal(fuelPercentage.Json())
	if err != nil {
		log.Error().AnErr("ProcessFuelPercentage: JSON Marshall", err).Send()
	} else {
		err := client.Publish(publishTopic, 0, false, dataToSend).Error()
		if err != nil {
			log.Error().AnErr("ProcessFuelPercentage: MQTT Publish", err).Send()
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
