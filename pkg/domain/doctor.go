package domain

import (
	"time"

	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
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

type DoctorExperience struct {
	ID              int64      `gorm:"primaryKey"`
	ProfileID       int64      `gorm:"not null"`
	Profile         UserDoctor `gorm:"foreignKey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	City            string     `gorm:"not null;size:100"`
	Province        string     `gorm:"not null;size:100"`
	InstitutionName string     `gorm:"not null;size:100"`
	AddressDetail   string     `gorm:"size:255"`
	StartDate       dtype.Date `gorm:"null;type:date"`
	EndDate         dtype.Date `gorm:"null;type:date"`
}

type DoctorEducation struct {
	ID         int64      `gorm:"primaryKey"`
	ProfileID  int64      `gorm:"not null"`
	Profile    UserDoctor `gorm:"foreignKey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	University string     `gorm:"not null;size:100"`
	Major      string     `gorm:"not null;size:100"`
	StartYear  int
	EndYear    int
}

type DoctorSchedule struct {
	ID        int64      `gorm:"primaryKey"`
	ProfileID int64      `gorm:"not null"`
	Profile   UserDoctor `gorm:"foreignKey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Date      dtype.Date `gorm:"not null;type:date;uniqueIndex:idx_doctor_schedules_profile_id_date"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`

	TimeSlots []DoctorScheduleTimeSlot `gorm:"foreignKey:ScheduleID;references:ID"`
}

type DoctorScheduleTimeSlot struct {
	ID         int64          `gorm:"primaryKey"`
	ScheduleID int64          `gorm:"not null"`
	Schedule   DoctorSchedule `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StartTime  string         `gorm:"size:5;not null"`
	EndTime    string         `gorm:"size:5;not null"`
	IsBooked   bool           `gorm:"not null;default:false"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
}
