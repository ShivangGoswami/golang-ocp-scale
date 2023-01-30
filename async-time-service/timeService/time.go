package time

import "time"

func GetTime() *time.Time {
	return &[]time.Time{time.Now()}[0]
}
