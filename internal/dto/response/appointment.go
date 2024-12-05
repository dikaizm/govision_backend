package response

import "github.com/dikaizm/govision_backend/pkg/helpers/dtype"

type GetAppointment struct {
	ID                   int64  `json:"id"`
	DoctorUserID         string `json:"doctor_user_id"`
	DoctorName           string `json:"doctor_name"`
	DoctorSpecialization string `json:"doctor_specialization"`
	DoctorPhotoURL       string `json:"doctor_photo"`

	Date      dtype.Date `json:"date"`
	StartTime string     `json:"start_time"`
	EndTime   string     `json:"end_time"`
	Location  string     `json:"location"`
}
