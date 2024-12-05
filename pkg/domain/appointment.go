package domain

import (
	"time"
)

type Appointment struct {
	ID         int64                  `gorm:"primaryKey"`
	PatientID  int64                  `gorm:"not null"`
	Patient    UserPatient            `gorm:"foreignKey:PatientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DoctorID   int64                  `gorm:"not null"`
	Doctor     UserDoctor             `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TimeSlotID int64                  `gorm:"not null;uniqueIndex:idx_appointments_time_slot_id"`
	TimeSlot   DoctorScheduleTimeSlot `gorm:"foreignKey:TimeSlotID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time              `gorm:"autoCreateTime"`
	UpdatedAt  time.Time              `gorm:"autoUpdateTime"`
}
