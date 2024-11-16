package repositories

import (
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

func (r *DbAppointmentRepository) Create(apt *domain.Appointment) error {
	if err := r.DB.Create(apt).Error; err != nil {
		return err
	}
	return nil
}

func (r *DbAppointmentRepository) FindAllByDoctor(doctorID int64) ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	if err := r.DB.Where("doctor_id = ?", doctorID).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *DbAppointmentRepository) FindAllByPatient(patientID int64) ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	if err := r.DB.Where("patient_id = ?", patientID).Find(&appointments).Error; err != nil {
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
