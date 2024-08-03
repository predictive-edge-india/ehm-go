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
	"github.com/predictive-edge-india/ehm-go/database"
	"github.com/predictive-edge-india/ehm-go/managers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/predictive-edge-india/ehm-go/routes"
	authRoutes "github.com/predictive-edge-india/ehm-go/routes/auth"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
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

func HandleErrors(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var e *fiber.Error
	log.Error().AnErr("Fiber error", err).Send()
	if errors.As(err, &e) {
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
	authRoutes.AuthRoutes(v1)

	routes.CustomerRoutes(v1)
	routes.UserRoutes(v1)

	routes.EhmDeviceRoutes(v1)
	routes.ParameterRoutes(v1)
}

func main() {
	config := fiber.Config{
		Prefork:      *PROD, // go run app.go -prod
		ErrorHandler: HandleErrors,
	}

	InitLogger()

	app = fiber.New(config)

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	database.InitDatabase()
	log.Info().Msg("Connected to DB!")

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", models.Environments.MqttHost, models.Environments.MqttPort))

	opts.SetClientID(uuid.NewString())
	opts.SetDefaultPublishHandler(f)
	topic := models.Environments.UniversalTopic

	opts.OnConnect = func(c MQTT.Client) {
		log.Info().Str("Subscribing", topic).Send()
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("Connected to MQTT server\n")
	}

	loadRoutes()

	envPort := models.Environments.ApiPort
	if len(envPort) == 0 {
		log.Fatal().AnErr("Failed starting service", app.Listen(HOST+":"+PORT)).Send()
	}
	log.Fatal().AnErr("Failed starting service", app.Listen(HOST+":"+envPort)).Send()
}

type SpecificLevelWriter struct {
	io.Writer
	Levels []zerolog.Level
}

func (w SpecificLevelWriter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	for _, l := range w.Levels {
		if l == level {
			return w.Write(p)
		}
	}
	return len(p), nil
}

func InitLogger() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.999Z07:00"}
	errorWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05.999Z07:00"}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(models.Environments.LogPath),
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}
	multiWriter := zerolog.MultiLevelWriter(
		SpecificLevelWriter{
			Writer: consoleWriter,
			Levels: []zerolog.Level{
				zerolog.DebugLevel, zerolog.InfoLevel, zerolog.WarnLevel,
			},
		},
		SpecificLevelWriter{
			Writer: errorWriter,
			Levels: []zerolog.Level{
				zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel,
			},
		},
		lumberjackLogger,
	)

	logger := zerolog.New(multiWriter).With().Timestamp().Logger()
	log.Logger = logger
}
