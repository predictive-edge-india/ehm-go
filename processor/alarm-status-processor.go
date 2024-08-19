package processor

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessAlarmStatus(client MQTT.Client, deviceId, message string) {
	device := database.FindDeviceBySerialNo(deviceId)
	if device.IsIdNull() {
		log.Error().Str("deviceId", deviceId).Send()
		return
	}

	flagsString := strings.Split(message, ",")
	flags := []float32{}

	for _, flag := range flagsString {
		flagFloat, err := strconv.ParseFloat(flag, 32)
		if err != nil {
			log.Error().AnErr("ProcessAlarmStatus: parse flag", err).Send()
			return
		}
		flags = append(flags, float32(flagFloat))
	}

	alarmStatusFlag := models.AlarmStatusFlag{
		DeviceId: device.Id,
		Statuses: flags,
	}

	err := database.Database.Create(&alarmStatusFlag).Error
	if err != nil {
		log.Error().AnErr("ProcessAlarmStatus: create alarmStatusFlag", err).Send()
		return
	}

	client.Publish("iisc/web/"+device.SerialNo+"/alarms", 0, false, alarmStatusFlag.MqttPayload())
}
