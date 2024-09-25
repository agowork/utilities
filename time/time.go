package time

import "time"

func GetNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowShortString() string {
	return time.Now().Format("20060102150405")
}

func GetTimeFromString(timeString string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timeString, time.Local)
}

func GetDaysAgoDDHHMMSS(days int) (string, string, string, string) {
	timestamp := time.Now().Add(-time.Hour * 24 * time.Duration(days))
	return timestamp.Format("2006-01-02"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}

func GetDaysAgoShortDDHHMMSS(days int) (string, string, string, string) {
	timestamp := time.Now().Add(-time.Hour * 24 * time.Duration(days))
	return timestamp.Format("20060102"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}
