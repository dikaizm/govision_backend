package helpers

import (
	"strconv"
	"time"
)

func StringToInt64(s string) (*int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}

func GetDaysOfWeek(start, end time.Time) ([]int, error) {
	daysOfWeekMap := make(map[int]bool)
	var daysOfWeek []int
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		day := int(d.Weekday())
		if day == 0 {
			day = 7
		}
		if !daysOfWeekMap[day] {
			daysOfWeekMap[day] = true
			daysOfWeek = append(daysOfWeek, day)
		}
	}

	return daysOfWeek, nil
}

func GetWorkYears(start time.Time, end time.Time) int {
	var year, month int
	year = end.Year() - start.Year()
	if year <= 1 {
		month = (int(end.Month()) + 12) - int(start.Month())
		if month <= 12 {
			year = 0
		}
	}

	return year
}
