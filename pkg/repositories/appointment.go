package repositories

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type DbAppointmentRepository struct {
	DB *gorm.DB
}

func NewDbAppointmentRepository(db *gorm.DB) repo_intf.AppointmentRepository {
	return &DbAppointmentRepository{DB: db}
}

func (r *DbAppointmentRepository) Create(apt *domain.Appointment) (*domain.Appointment, error) {
	if err := r.DB.Create(apt).Error; err != nil {
		return nil, err
	}
	return apt, nil
}

func (r *DbAppointmentRepository) FindAllByDoctor(doctorID int64) ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	if err := r.DB.Where("doctor_id = ?", doctorID).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *DbAppointmentRepository) FindAllByPatient(patientID int64, filter *request.FilterViewAllAppointment) ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment

	// Start the query with Preload statements for related entities
	query := r.DB.Preload("Patient").Preload("Doctor").Preload("Doctor.User").Preload("TimeSlot.Schedule").Where("patient_id = ?", patientID)

	if filter != nil {
		switch filter.Range {
		case "today":
			query = query.Joins("JOIN doctor_schedule_time_slots ON appointments.time_slot_id = doctor_schedule_time_slots.id").
				Joins("JOIN doctor_schedules ON doctor_schedule_time_slots.schedule_id = doctor_schedules.id").
				Where("doctor_schedules.date = CURRENT_DATE")
		case "week":
			query = query.Joins("JOIN doctor_schedule_time_slots ON appointments.time_slot_id = doctor_schedule_time_slots.id").
				Joins("JOIN doctor_schedules ON doctor_schedule_time_slots.schedule_id = doctor_schedules.id").
				Where("DATE_PART('year', doctor_schedules.date) = DATE_PART('year', CURRENT_DATE) AND " +
					"DATE_PART('week', doctor_schedules.date) = DATE_PART('week', CURRENT_DATE)")
		case "month":
			query = query.Joins("JOIN doctor_schedule_time_slots ON appointments.time_slot_id = doctor_schedule_time_slots.id").
				Joins("JOIN doctor_schedules ON doctor_schedule_time_slots.schedule_id = doctor_schedules.id").
				Where("DATE_PART('year', doctor_schedules.date) = DATE_PART('year', CURRENT_DATE) AND " +
					"DATE_PART('month', doctor_schedules.date) = DATE_PART('month', CURRENT_DATE)")
		}
	}

	// Execute the query
	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

func (r *DbAppointmentRepository) UpdateStatus(id int64, status string) error {
	if err := r.DB.Model(&domain.Appointment{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *DbAppointmentRepository) Delete(id int64) error {
	if err := r.DB.Where("id = ?", id).Delete(&domain.Appointment{}).Error; err != nil {
		return err
	}
	return nil
}
