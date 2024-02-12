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
