// Create by Yale 2019/1/10 16:34
package ctime

import "time"

const (
	TIME_TEMP="2006-01-02 15:04:05"
	DATE_TEMP="2006-01-02"
)

func CheckDate(dateStr string)(time.Time,error) {
	return time.Parse(DATE_TEMP,dateStr)
}
func GetDateStr(t time.Time) string {
	return t.Format(DATE_TEMP)
}

