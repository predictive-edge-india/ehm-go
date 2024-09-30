package processor

import (
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
)

func ProcessGps(client MQTT.Client, deviceId, message string) {
	device := database.FindDeviceBySerialNo(deviceId)
	if device.IsIdNull() {
		log.Error().Str("deviceId", deviceId).Send()
		return
	}

	gpsSplitStr := strings.Split(message, ",")
	if len(gpsSplitStr) != 2 {
		log.Error().Str("deviceId", deviceId).Str("message", message).Send()
		return
	}

	gpsLat := helpers.ParseFloat64(gpsSplitStr[0])
	gpsLng := helpers.ParseFloat64(gpsSplitStr[1])

	if gpsLat <= 0 || gpsLng <= 0 {
		log.Error().Str("deviceId", deviceId).Float64("gpsLat", gpsLat).Float64("gpsLng", gpsLng).Send()
		return
	}

	deviceLastLocation := models.DeviceLastLocation{
		DeviceId: device.Id,
		Position: models.GeoJson{Type: "Point", Coordinates: []float64{gpsLng, gpsLat}},
	}

	err := database.Database.Create(&deviceLastLocation).Error
	if err != nil {
		log.Error().AnErr("ProcessGps: create deviceLastLocation", err).Send()
		return
	}

	client.Publish("iisc/web/"+device.SerialNo+"/gps", 0, false, deviceLastLocation.MqttPayload())
}
