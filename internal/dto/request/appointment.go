package request

import (
	"time"

	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
)

type (
	CreateAppointment struct {
		PatientUserID string
		DoctorUserID  string     `json:"doctor_user_id" validate:"required"`
		Date          dtype.Date `json:"date" validate:"required"`
		TimeSlotID    int64      `json:"time_slot_id" validate:"required"`
	}

	ViewAppointment struct {
		UserID   string
		UserRole string
	}

	FilterViewAllAppointment struct {
		Range string `schema:"range"` // today, week, month
	}

	ConfirmAppointment struct {
		Confirm bool `json:"confirm" validate:"required"`
	}

	FilterAppointmentSchedule struct {
		Days      []string  `schema:"days"`
		StartDate time.Time `schema:"start_date"`
		EndDate   time.Time `schema:"end_date"`
		StartHour string    `schema:"start_hour"`
		EndHour   string    `schema:"end_hour"`
	}
)
