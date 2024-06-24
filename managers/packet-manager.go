package managers

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/iisc/demo-go/helpers"
	"github.com/iisc/demo-go/processor"
)

func ProcessPacket(client MQTT.Client, topic, message string) {
	topicType := helpers.GetTopicType(topic)
	if topicType == 1 {
		processor.ProcessCurrentMessage(client, topic, message)
	} else if topicType == 2 {
		processor.ProcessFuelPercentage(client, topic, message)
	} else if topicType == 3 {
		processor.ProcessFaults(client, topic, message)
	} else if topicType == 4 {
		processor.ProcessPowerParam(client, topic, message)
	} else if topicType == 5 {
		processor.ProcessEngineParam(client, topic, message)
	}
}
