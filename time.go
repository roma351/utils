package utils

import "time"

func GetDateString(value time.Time) string {
	return value.Format("2006-01-02")
}
