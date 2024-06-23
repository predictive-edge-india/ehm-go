package managers

import (
	"strconv"
	"strings"

	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/helpers"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
)

func ProcessPacket(topic, message string) *models.CurrentParameter {
	newParam := new(models.CurrentParameter)
	deviceId, paramType, dataType, err := helpers.GetTopicData(topic)
	if err != nil {
		log.Errorln(err.Error())
		return newParam
	}

	ehmDevice, err := database.FindOrCreateEhmDevice(deviceId)
	if err != nil {
		log.Errorln(err.Error())
		return newParam
	}

	newParam.ParamType = paramType
	newParam.EhmDeviceId = &ehmDevice.Id

	if dataType == models.InputDataType.RMS {
		rmsValue, err := parseRms(message)
		if err != nil {
			log.Errorln(err.Error())
			return newParam
		}
		newParam.RMS = helpers.ToFixed(rmsValue, 2)
	} else if dataType == models.InputDataType.FFT {
		fftValues := ParseFft(message)
		newParam.FFT = fftValues
	}

	newParam.PacketType = dataType
	return newParam
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
