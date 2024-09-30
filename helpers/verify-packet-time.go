package helpers

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func GetPayloadTime(payload string) (time.Time, error) {
	timeSplit := strings.Split(payload, ":")
	if len(timeSplit) != 2 {
		return time.Now(), errors.New("packet format error")
	}

	timeStr := timeSplit[1]
	unixTime, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	timeStamp := time.Unix(unixTime, 0)
	if timeStamp.After(time.Now()) {
		return timeStamp, errors.New("invalid timestamp")
	}

	return timeStamp, nil
}
