package managers

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/iisc/demo-go/helpers"
)

func ProcessPacket(client MQTT.Client, topic, message string) {
	topicType := helpers.GetTopicType(topic)
	if topicType == 1 {
		ProcessCurrentMessage(client, topic, message)
	} else if topicType == 2 {
		ProcessFuelPercentage(client, topic, message)
	}
}

func parseRms(message string) (float64, error) {
	rmsValue, err := strconv.ParseFloat(message, 64)
	return float64(rmsValue), err
}

func ParseFft(message string) [][]float64 {
	var fftValues [][]float64

	cleanedMessage := strings.ReplaceAll(message, "\n", "")

	plots := strings.Split(cleanedMessage, ":")

	for _, plot := range plots {
		coords := strings.Split(plot, ",")
		if len(coords) == 2 && helpers.ParseFloat64(coords[0]) > 0 {
			x := helpers.ParseFloat64(coords[0])
			y := helpers.ParseFloat64(coords[1])
			p := []float64{x, y}
			fftValues = append(fftValues, p)
		}
	}

	if len(fftValues) > 0 {
		fftValues = fftValues[1:]
	}

	return fftValues
}
