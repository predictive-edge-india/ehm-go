package processor

import (
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessDGStatus(client MQTT.Client, serialNo, message string) {
	device := database.FindDeviceBySerialNo(serialNo)
	if device.IsIdNull() {
		log.Error().Str("ProcessDGStatus: Device not found with serial no", serialNo).Send()
		return
	}

	dgStatusesStr := strings.Split(message, ",")
	dgStatuses := []uint16{}

	for _, dgStatus := range dgStatusesStr {
		statusInt16, err := strconv.ParseInt(dgStatus, 10, 16)
		if err != nil {
			log.Error().AnErr("ProcessDGStatus: Parse ", err).Send()
			return
		}
		dgStatuses = append(dgStatuses, uint16(statusInt16))
	}

	if len(dgStatuses) != 21 {
		log.Error().Str("ProcessDGStatus: Wrong dgStatuses", serialNo).Send()
		return
	}

	newDgStatus := models.DGStatus{
		DeviceId: device.Id,
	}

	newDgStatus.SGCMonitoringMode = dgStatuses[0] == 1
	newDgStatus.MainsHealthy = dgStatuses[1] == 1

	newDgStatus.DGOperationMode = uint8(dgStatuses[2])

	newDgStatus.LoadOnMains = dgStatuses[3] == 1
	newDgStatus.LoadOnDG = dgStatuses[4] == 1
	newDgStatus.CurrentDGStatus = dgStatuses[5] == 1
	newDgStatus.DGStoppedNormally = dgStatuses[6] == 1
	newDgStatus.DGStoppedWithFault = dgStatuses[7] == 1
	newDgStatus.DGFailToStart = dgStatuses[8] == 1
	newDgStatus.GenAvailable = dgStatuses[9] == 1
	newDgStatus.CommonShutDown = dgStatuses[10] == 1
	newDgStatus.CommonElectricTrip = dgStatuses[11] == 1
	newDgStatus.CommonWarning = dgStatuses[12] == 1
	newDgStatus.CommonNotification = dgStatuses[13] == 1

	newDgStatus.CurrentTimeStampMin = uint8(dgStatuses[14])
	newDgStatus.CurrentTimeStampSec = uint8(dgStatuses[15])
	newDgStatus.CurrentTimeStampWeekDay = uint8(dgStatuses[16])
	newDgStatus.CurrentTimeStampHour = uint8(dgStatuses[17])
	newDgStatus.CurrentTimeStampMonth = uint8(dgStatuses[18])
	newDgStatus.CurrentTimeStampDay = uint8(dgStatuses[19])
	newDgStatus.CurrentTimeStampYear = uint16(dgStatuses[20])

	if newDgStatus.GetUnixTimestamp() <= 0 {
		log.Error().Str("ProcessDGStatus: Wrong timestamp", serialNo).Send()
		return
	}

	err := database.Database.Create(&newDgStatus).Error
	if err != nil {
		log.Error().AnErr("ProcessDGStatus: create newDgStatus", err).Send()
		return
	}

	client.Publish("iisc/web/"+device.SerialNo+"/dgstatus", 0, false, newDgStatus.MqttPayload())
}
