package util

import (
	"time"
)

func ThisWeekRange() []time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	t := time.Now().In(loc)

	//Back to this day last week
	t = t.AddDate(0, 0, -7)

	days := []time.Time{}
	for i := 0; i < 7; i++ {
		days = append(days, t)
		t = t.AddDate(0, 0, 1)
	}

	return days
}
