package utils

import "time"

func GetToday() string {
	timezone := "Asia/Bangkok"
	location, _ := time.LoadLocation(timezone)
	return time.Now().In(location).Format("02")
}
func GetCurrentDate() string {
	timezone := "Asia/Bangkok"
	location, _ := time.LoadLocation(timezone)
	return time.Now().In(location).Format("2006-01-02")
}

func GetCurrentMonth() string {
	timezone := "Asia/Bangkok"
	location, _ := time.LoadLocation(timezone)
	return time.Now().In(location).Format("2006-01")
}
