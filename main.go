package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/iisc/demo-go/database"
	"github.com/iisc/demo-go/managers"
	"github.com/iisc/demo-go/models"
	"github.com/iisc/demo-go/routes"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	HOST = "127.0.0.1"
	PORT = "4000"
)

var (
	PROD = flag.Bool("prod", false, "Enable prefork in Production")
	app  *fiber.App
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("%s: %s\n", msg.Topic(), msg.Payload())
	managers.ProcessPacket(client, msg.Topic(), string(msg.Payload()))
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

func HandleErrors(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var e *fiber.Error
	if errors.As(err, &e) {
		log.Println(e.Message)
		code = e.Code
		message = e.Message
	}

	err = ctx.Status(code).JSON(fiber.Map{
		"type":    "error",
		"message": message,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return nil
}

func loadRoutes() {
	v1 := app.Group("/v1")

	routes.EhmDeviceRoutes(v1)
}

func main() {
	config := fiber.Config{
		Prefork:      *PROD, // go run app.go -prod
		ErrorHandler: HandleErrors,
	}

	initLogger()

	app = fiber.New(config)

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	database.InitDatabase()
	log.Info("Connected to DB!")

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", models.Environments.MqttHost, models.Environments.MqttPort))

	opts.SetClientID(uuid.NewString())
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

	loadRoutes()

	envPort := models.Environments.ApiPort
	if len(envPort) == 0 {
		log.Fatal(app.Listen(HOST + ":" + PORT))
	}
	log.Fatal(app.Listen(HOST + ":" + envPort))
}
