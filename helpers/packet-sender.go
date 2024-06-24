package helpers

import (
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	log "github.com/sirupsen/logrus"
)

func PublishToTopic(client MQTT.Client, topic string, data map[string]interface{}) {
	log.Infoln("Publishing to topic: " + topic)
	dataToSend, err := json.Marshal(data)
	if err != nil {
		log.Errorln(err.Error())
	} else {
		err := client.Publish(topic, 0, false, dataToSend).Error()
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}
