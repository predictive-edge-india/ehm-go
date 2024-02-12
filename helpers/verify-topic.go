package helpers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/iisc/demo-go/models"
)

func GetTopicData(topic string) (string, string, int32, error) {
	pattern := `^iisc/ehm/.*\/(i1|i2|i3)/(rms|fft)$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(topic) {
		return "", "", -1, errors.New("invalid topic data 1")
	}

	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 3 {
		return "", "", -1, errors.New("invalid topic data 2")
	}

	if rawStringArr[2] == "rms" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.RMS, nil
	}
	if rawStringArr[2] == "fft" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.FFT, nil
	}

	return "", "", -1, errors.New("invalid topic type 3")
}
