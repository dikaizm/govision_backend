package response

import (
	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
)

type GetDoctorProfilePreview struct {
	UserID         string  `json:"user_id"`
	Name           string  `json:"name"`
	Specialization string  `json:"specialization"`
	WorkYears      int     `json:"work_years"`
	Rating         float64 `json:"rating"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Photo          string  `json:"photo"`
}

type GetDoctorProfile struct {
	UserID         string             `json:"user_id"`
	Name           string             `json:"name"`
	StrNo          string             `json:"str_no"`
	Photo          string             `json:"photo"`
	Specialization string             `json:"specialization"`
	Institution    string             `json:"institution"`
	City           string             `json:"city"`
	Province       string             `json:"province"`
	Rating         float64            `json:"rating"`
	TotalPatient   int                `json:"total_patient"`
	WorkYears      int                `json:"work_years"`
	BioDesc        string             `json:"bio_desc"`
	Experiences    []DoctorExperience `json:"experiences"`
	Educations     []DoctorEducation  `json:"educations"`
	Schedules      []DoctorSchedule   `json:"schedules"`
}

type DoctorExperience struct {
	Institution string     `json:"institution"`
	StartDate   dtype.Date `json:"start_date"`
	EndDate     dtype.Date `json:"end_date"`
}

type DoctorEducation struct {
	University string `json:"university"`
	Major      string `json:"major"`
	StartYear  int    `json:"start_year"`
	EndYear    int    `json:"end_year"`
}

type DoctorSchedule struct {
	Date dtype.Date `json:"date"`
}

type GetDoctorTimeSlot struct {
	ID        int64  `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	IsBooked  bool   `json:"is_booked"`
}
