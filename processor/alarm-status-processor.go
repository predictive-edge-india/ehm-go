package processor

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessAlarmStatus(client MQTT.Client, serialNo, message string) {
	device := database.FindDeviceBySerialNo(serialNo)
	if device.IsIdNull() {
		log.Error().Str("ProcessAlarmStatus: Device not found with serial no", serialNo).Send()
		return
	}

	flagsString := strings.Split(message, ",")
	flags := []uint8{}

	for _, flag := range flagsString {
		flagInt8, err := strconv.ParseInt(flag, 10, 8)
		if err != nil {
			log.Error().AnErr("ProcessAlarmStatus: Parse ", err).Send()
			return
		}
		flags = append(flags, uint8(flagInt8))
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
