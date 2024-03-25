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

// TimeCalendar  根据开始时间和结束时间生成期间每天的开始时间和结束时间
func TimeCalendar(start, end time.Time, dayCalendarCallback func(day string, start, end time.Time)) {
	for d := start; d.Unix() <= end.Unix(); d = d.AddDate(0, 0, 1) {
		dailyStart := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
		dailyEnd := dailyStart.Add(24*60*60 - 1*time.Second)
		dayCalendarCallback(d.Format("2006-01-02"), dailyStart, dailyEnd)
	}
}
