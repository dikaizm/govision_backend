package request

import "time"

type (
	CreateAppointment struct {
		PatientID int64     `json:"patient_id" validate:"required"`
		DoctorID  int64     `json:"doctor_id" validate:"required"`
		Date      time.Time `json:"date" validate:"required"`
		StartHour time.Time `json:"start_hour" validate:"required"`
		EndHour   time.Time `json:"end_hour" validate:"required"`
	}

	ViewAppointment struct {
		UserID   string
		UserRole string
	}

	ConfirmAppointment struct {
		Confirm bool `json:"confirm" validate:"required"`
	}

	FilterAppointmentSchedule struct {
		DaysInt   []int     `json:"days_int"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		StartHour string    `json:"start_hour"`
		EndHour   string    `json:"end_hour"`
	}
)
