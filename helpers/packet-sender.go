package helpers

import (
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

func PublishToTopic(client MQTT.Client, topic string, data map[string]interface{}) {
	log.Info().Str("Publishing topic", topic).Send()
	dataToSend, err := json.Marshal(data)
	if err != nil {
		log.Error().AnErr("PublishToTopic: JSON Marshall", err).Send()
	} else {
		if err := client.Publish(topic, 0, false, dataToSend).Error(); err != nil {
			log.Error().AnErr("PublishToTopic publish error", err).Send()
		}
	}
}
