package calc

import (
	"time"
)

// WeekCount 计算两个时间之间的周数差值
// start 开始时间
// end 结束时间
func WeekCount(start, end time.Time) int {
	currentWeek := int(start.Weekday())
	if currentWeek == 0 {
		currentWeek = 7
	}
	endWeek := int(end.Weekday())
	if endWeek == 0 {
		endWeek = 7
	}
	if endWeek != currentWeek {
		end = end.AddDate(0, 0, currentWeek-endWeek)
	}
	return int(end.Sub(start).Hours() / 24 / 7)
}
