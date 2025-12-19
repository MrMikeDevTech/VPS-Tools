package utils

import (
	"fmt"
	"time"
)

var uptimeObj struct {
	Days    int64
	Hours   int64
	Minutes int64
	Seconds int64
}

func UptimeFormat(uptRaw time.Duration) string {
	totalSeconds := int64(uptRaw.Seconds())
	uptimeObj.Days = totalSeconds / 86400
	totalSeconds %= 86400
	uptimeObj.Hours = totalSeconds / 3600
	totalSeconds %= 3600
	uptimeObj.Minutes = totalSeconds / 60
	uptimeObj.Seconds = totalSeconds % 60

	uptimeStr := ""

	if uptimeObj.Days > 0 {
		if uptimeObj.Days > 1 {
			uptimeStr += fmt.Sprintf("%d días ", uptimeObj.Days)
		} else {
			uptimeStr += fmt.Sprintf("%d día ", uptimeObj.Days)
		}
	}

	if uptimeObj.Hours > 0 {
		if uptimeObj.Hours > 1 {
			uptimeStr += fmt.Sprintf("%d horas ", uptimeObj.Hours)
		} else {
			uptimeStr += fmt.Sprintf("%d hora ", uptimeObj.Hours)
		}
	}

	if uptimeObj.Minutes > 0 {
		if uptimeObj.Minutes > 1 {
			uptimeStr += fmt.Sprintf("%d minutos ", uptimeObj.Minutes)
		} else {
			uptimeStr += fmt.Sprintf("%d minuto ", uptimeObj.Minutes)
		}
	}

	if uptimeObj.Seconds > 0 {
		if uptimeObj.Seconds > 1 {
			uptimeStr += fmt.Sprintf("%d segundos", uptimeObj.Seconds)
		} else {
			uptimeStr += fmt.Sprintf("%d segundo", uptimeObj.Seconds)
		}
	}

	return uptimeStr
}
