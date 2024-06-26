package helpers

import (
	"math"
	"strconv"
)

func ParseFloat64(str string) float64 {
	newVal, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return -1
	}
	return float64(newVal)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := float64(math.Pow(10, float64(precision)))
	return float64(round(num*output)) / output
}

func StringToUint16(str string) uint16 {
	newVal, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(newVal)
}

func StringToUint8(str string) uint8 {
	newVal, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(newVal)
}

func ParseFloat32(str string) float32 {
	newVal, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(newVal)
}

func StringToInt8(str string) int8 {
	newVal, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0
	}
	return int8(newVal)
}

func KelvinToCelsius(kelvin float32) float32 {
	return kelvin - 273.15
}

func Float64ToFloat32(f float64) float32 {
	return float32(f)
}
