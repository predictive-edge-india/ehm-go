package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/managers"
	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var gRms float64

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("%s: %s\n", msg.Topic(), msg.Payload())
	currentParam := managers.ProcessPacket(msg.Topic(), string(msg.Payload()))

	if currentParam.PacketType == models.InputDataType.RMS {
		gRms = currentParam.RMS
	}
	if currentParam.PacketType == models.InputDataType.FFT {
		currentParam.RMS = gRms
		database.CreateCurrentParameter(*currentParam)
	}
}

func initLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(models.Environments.LogPath),
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}
	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	logFormatter := new(log.TextFormatter)
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)
	log.SetOutput(multiWriter)
}

func main() {
	initLogger()
	database.InitDatabase()
	log.Info("Connected to DB!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", models.Environments.TcpIp, models.Environments.TcpPort))

	opts.SetClientID("magnus-mbp")
	opts.SetDefaultPublishHandler(f)
	topic := models.Environments.UniversalTopic

	opts.OnConnect = func(c MQTT.Client) {
		log.Println("Subscribing to universal topic: ", topic)
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}
	<-c
}
