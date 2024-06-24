package processor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/helpers"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

var gI1Rms float64
var gXRms float64

func ProcessCurrentMessage(client MQTT.Client, topic string, message string) {
	newParam := new(models.CurrentParameter)

	deviceId, paramType, dataType, err := processCurrentTopic(topic)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	newParam.ParamType = paramType
	newParam.EhmDeviceId = &ehmDevice.Id

	if dataType == models.InputDataType.RMS {
		rmsValue, err := parseRms(message)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		newParam.RMS = helpers.ToFixed(rmsValue, 2)
	} else if dataType == models.InputDataType.FFT {
		fftValues := ParseFft(message)
		newParam.FFT = fftValues
	}

	newParam.PacketType = dataType

	if newParam.PacketType == models.InputDataType.RMS {
		if paramType == "i1" {
			gI1Rms = newParam.RMS
		} else if paramType == "x" {
			gXRms = newParam.RMS
		}
	}

	if newParam.PacketType == models.InputDataType.FFT {
		if gI1Rms > 0 && paramType == "i1" {
			newParam.RMS = gI1Rms
			newParam = database.CreateCurrentParameter(*newParam)
		} else if gXRms > 0 && paramType == "x" {
			newParam.RMS = gXRms
			newParam = database.CreateCurrentParameter(*newParam)
		}

		publishTopic := fmt.Sprintf("iisc/web/%s/current/fft/%s", newParam.EhmDeviceId, newParam.ParamType)
		helpers.PublishToTopic(client, publishTopic, newParam.Json())
	}
}

func processCurrentTopic(topic string) (string, string, int32, error) {

	rawString := strings.Replace(topic, "iisc/ehm/", "", 1)
	rawStringArr := strings.Split(rawString, "/")

	if len(rawStringArr) != 3 {
		return "", "", -1, errors.New("current topic invalid length")
	}

	if rawStringArr[2] == "rms" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.RMS, nil
	}
	if rawStringArr[2] == "fft" {
		return rawStringArr[0], rawStringArr[1], models.InputDataType.FFT, nil
	}

	return "", "", -1, errors.New("current topic invalid type")
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
