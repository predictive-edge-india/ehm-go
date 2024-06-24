package helpers

import (
	"regexp"
)

func GetTopicType(topic string) int32 {
	// current topic
	pattern := `^iisc/ehm/.*\/(i1|i2|i3)/(rms|fft)$`
	regex := regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return 1
	}

	// fuel percentage topic
	pattern = `^iisc\/ehm\/\d+\/fuel_p$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return 2
	}

	// fault topic
	pattern = `^iisc\/ehm\/\d+\/faults$`
	regex = regexp.MustCompile(pattern)
	if regex.MatchString(topic) {
		return 3
	}

	return -1
}
