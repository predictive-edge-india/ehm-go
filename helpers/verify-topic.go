package helpers

import (
	"regexp"
	"strings"
)

func GetTopicType(topic string) (string, int8) {
	// current topic
	splitTopic := strings.Split(topic, "/")

	deviceId := splitTopic[2]
	if len(splitTopic) > 3 && splitTopic[0] == "iisc" && splitTopic[1] == "ehm" {
		deviceId = splitTopic[2]
	}

	pattern := `^iisc/ehm/.*\/gps$`
	regex := regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 1
	}

	pattern = `^iisc/ehm/.*\/power$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 2
	}

	pattern = `^iisc/ehm/.*\/alarms$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 3
	}

	pattern = `^iisc/ehm/.*\/dgStatus$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 4
	}

	pattern = `^iisc/ehm/.*\/e483Can$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 5
	}

	pattern = `^iisc/ehm/.*\/deviceTime$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return deviceId, 6
	}

	return deviceId, -1
}
