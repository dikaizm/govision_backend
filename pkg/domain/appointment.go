package domain

import (
	"time"
)

const (
	AppointmentStatusPending     = "pending"
	AppointmentStatusAccepted    = "accepted"
	AppointmentStatusRejected    = "rejected"
	AppointmentStatusCanceled    = "canceled"
	AppointmentStatusRescheduled = "rescheduled"
)

type Appointment struct {
	ID        int64       `gorm:"primaryKey"`
	PatientID int64       `gorm:"not null"`
	Patient   UserPatient `gorm:"foreignKey:PatientID"`
	DoctorID  int64       `gorm:"not null"`
	Doctor    UserDoctor  `gorm:"foreignKey:DoctorID"`
	Date      time.Time
	StartHour time.Time
	EndHour   time.Time
	Status    string `gorm:"size:20;check:status IN ('pending','accepted','rejected','canceled','rescheduled')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
