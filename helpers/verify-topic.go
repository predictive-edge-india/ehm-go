package helpers

import (
	"regexp"
	"strings"
)

func GetTopicType(topic string) (string, int16) {
	// current topic
	pattern := `^iisc/ehm/.*\/gps$`
	splitTopic := strings.Split(topic, "/")

	deviceId := splitTopic[2]
	if len(splitTopic) > 3 && splitTopic[0] == "iisc" && splitTopic[1] == "ehm" {
		deviceId = splitTopic[2]
	}

	regex := regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 1
	}

	return deviceId, -1
}
