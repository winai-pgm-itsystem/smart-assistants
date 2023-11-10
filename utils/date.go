package utils

import "time"

func GetCurrentDate(timezone string) string {
	location, _ := time.LoadLocation(timezone)
	return time.Now().In(location).Format("2006-01-02")
}
