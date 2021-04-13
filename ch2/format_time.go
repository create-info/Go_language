package ch2

import "time"

const MSTTimeFormatZ = "2006-01-02T15:04:05.000Z07:00"

func GetDateTimeWithTimeZone() string {
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
	return now.Format(MSTTimeFormatZ)
	}
	return now.In(loc).Format(MSTTimeFormatZ)
}