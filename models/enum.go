package models

import "os"

var Environments = newEnvironmentRegistry()

func newEnvironmentRegistry() *environmentRegistry {
	return &environmentRegistry{
		DBHost:         os.Getenv("IISC_DB_HOST"),
		DBUser:         os.Getenv("IISC_DB_USER"),
		DBPassword:     os.Getenv("IISC_DB_PASSWORD"),
		DBName:         os.Getenv("IISC_DB_NAME"),
		DBPort:         os.Getenv("IISC_DB_PORT"),
		ApiPort:        os.Getenv("API_PORT"),
		LogPath:        os.Getenv("LOG_PATH"),
		TcpIp:          os.Getenv("TCP_IP"),
		TcpPort:        os.Getenv("TCP_PORT"),
		UniversalTopic: os.Getenv("UNIVERSAL_TOPIC"),
	}
}

type environmentRegistry struct {
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	ApiPort        string
	LogPath        string
	TcpIp          string
	TcpPort        string
	UniversalTopic string
}

var InputDataType = newInputDataTypeRegistry()

func newInputDataTypeRegistry() *otpTypeRegistry {
	return &otpTypeRegistry{
		RMS: 1,
		FFT: 2,
	}
}

type otpTypeRegistry struct {
	RMS int32
	FFT int32
}
