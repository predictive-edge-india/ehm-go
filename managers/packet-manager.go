package managers

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/processor"
)

func ProcessPacket(client MQTT.Client, topic, message string) {
	deviceId, topicType := helpers.GetTopicType(topic)
	if topicType == 1 {
		processor.ProcessGps(client, deviceId, message)
	}
}
