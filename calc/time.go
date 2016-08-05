package calc

import (
	"time"
)

// WeekCount calculate the number of weeks between the two time difference
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
