package app

import "time"

func parseYMDToTime(ymd string) (time.Time, error) {
	return time.Parse("2006-01-02", ymd)
}
