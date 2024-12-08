package helpers

import (
	"strconv"
	"time"

	"github.com/dikaizm/govision_backend/pkg/domain"
	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
)

func StringToInt64(s string) (*int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
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

func GetWorkYears(start dtype.Date, end dtype.Date) int {
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

func GetRoleIDByName(roleName string, roles []*domain.UserRole) int {
	for _, role := range roles {
		if role.RoleName == roleName {
			return role.ID
		}
	}
	return -1
}
