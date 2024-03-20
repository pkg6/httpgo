package t

import (
	"time"
)

// TimeBetween 判断开始时间和结束时间是否小于等于diff
func TimeBetween(start, end time.Time, diff time.Duration) bool {
	sub := end.Sub(start)
	if sub.Microseconds() < 0 {
		return false
	}
	return sub <= diff
}
