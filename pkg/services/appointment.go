package services

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type AppointmentService struct {
	aptRepo    repo_intf.AppointmentRepository
	userRepo   repo_intf.UserRepository
	doctorRepo repo_intf.DoctorRepository
}

func NewAppointmentService(aptRepo repo_intf.AppointmentRepository, userRepo repo_intf.UserRepository, doctorRepo repo_intf.DoctorRepository) service_intf.AppointmentService {
	return &AppointmentService{
		aptRepo:    aptRepo,
		userRepo:   userRepo,
		doctorRepo: doctorRepo,
	}
}

func (u *AppointmentService) Create(p *request.CreateAppointment) (*domain.Appointment, error) {
	doctor, err := u.userRepo.FindDoctorProfileByID(p.DoctorUserID)
	if err != nil {
		return nil, err
	}

	patient, err := u.userRepo.FindPatientProfileByID(p.PatientUserID)
	if err != nil {
		return nil, err
	}

	apt := &domain.Appointment{
		PatientID:  patient.ID,
		DoctorID:   doctor.ID,
		TimeSlotID: p.TimeSlotID,
	}

	newApt, err := u.aptRepo.Create(apt)
	if err != nil {
		return nil, err
	}

	// Change time slot status to booked
	if err := u.doctorRepo.UpdateTimeSlotToBooked(p.TimeSlotID); err != nil {
		return nil, err
	}

	return newApt, nil
}

func (u *AppointmentService) FindAllByDoctor(userID string) ([]*domain.Appointment, error) {
	doctor, err := u.userRepo.FindDoctorProfileByID(userID)
	if err != nil {
		return nil, err
	}

	appointments, err := u.aptRepo.FindAllByDoctor(doctor.ID)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (u *AppointmentService) FindAllByPatient(userID string, filter *request.FilterViewAllAppointment) ([]*domain.Appointment, error) {
	patient, err := u.userRepo.FindPatientProfileByID(userID)
	if err != nil {
		return nil, err
	}

	appointments, err := u.aptRepo.FindAllByPatient(patient.ID, filter)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (u *AppointmentService) UpdateStatus(aptID int64, confirm bool) error {
	var status string
	if confirm {
		status = "confirmed"
	} else {
		status = "rejected"
	}

	if err := u.aptRepo.UpdateStatus(aptID, status); err != nil {
		return err
	}

	return nil
}
