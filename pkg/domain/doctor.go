package domain

import (
	"time"
)

const (
	DayMonday    = "monday"
	DayTuesday   = "tuesday"
	DayWednesday = "wednesday"
	DayThursday  = "thursday"
	DayFriday    = "friday"
	DaySaturday  = "saturday"
	DaySunday    = "sunday"
)

type DoctorPractice struct {
	ID            int64      `gorm:"primaryKey"`
	ProfileID     int64      `gorm:"not null"`
	Profile       UserDoctor `gorm:"foreignKey:ProfileID"`
	City          string     `gorm:"not null;size:100"`
	Province      string     `gorm:"not null;size:100"`
	OfficeName    string     `gorm:"not null;size:100"`
	AddressDetail string     `gorm:"not null;size:255"`
	StartDate     time.Time
	EndDate       time.Time
}

type DoctorEducation struct {
	ID         int64      `gorm:"primaryKey"`
	ProfileID  int64      `gorm:"not null"`
	Profile    UserDoctor `gorm:"foreignKey:ProfileID"`
	University string     `gorm:"not null;size:100"`
	Major      string     `gorm:"not null;size:100"`
	StartYear  int
	EndYear    int
}

type DoctorSchedule struct {
	ID        int64      `gorm:"primaryKey"`
	ProfileID int64      `gorm:"not null"`
	Profile   UserDoctor `gorm:"foreignKey:ProfileID"`
	DayOfWeek string     `gorm:"size:10;check:day_of_week IN ('monday','tuesday','wednesday','thursday','friday','saturday','sunday')"`
	StartHour time.Time
	EndHour   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
