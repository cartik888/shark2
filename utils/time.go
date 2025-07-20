package utils

import (
	"time"
)

func GetCurrentTime() time.Time {
	return time.Now()
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
