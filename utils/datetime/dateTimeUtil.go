package datetime

import "time"

const (
	STANDRD_DAY  = "2006-01-02"
	STANDRD_TIME = "2006-01-02 15:04:05"
)

func TimeToString(thisTime time.Time, format string) string {
	return thisTime.Format(format)
}
