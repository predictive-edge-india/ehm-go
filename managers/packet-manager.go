package managers

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"

	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/processor"
)

func ProcessPacket(client MQTT.Client, topic, message string) {
	deviceId, topicType := helpers.GetTopicType(topic)

	packetTime, err := helpers.GetPayloadTime(message)
	if err != nil {
		log.Error().AnErr("ProcessPacket: GetPayloadTime", err).Send()
		packetTime = time.Now()
	}

	log.Info().Int8("Topic", topicType).Str("DeviceId", deviceId).Send()

	if topicType == 1 {
		processor.ProcessGps(client, deviceId, message, packetTime)
	} else if topicType == 3 {
		processor.ProcessAlarmStatus(client, deviceId, message)
	} else if topicType == 4 {
		processor.ProcessDGStatus(client, deviceId, message)
	} else if topicType == 2 {
		processor.ProcessPowerData(client, deviceId, message)
	}
}
