package services

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type AppointmentService struct {
	aptRepo  repo_intf.AppointmentRepository
	userRepo repo_intf.UserRepository
}

func NewAppointmentService(aptRepo repo_intf.AppointmentRepository, userRepo repo_intf.UserRepository) service_intf.AppointmentService {
	return &AppointmentService{
		aptRepo:  aptRepo,
		userRepo: userRepo,
	}
}

func (u *AppointmentService) Create(p *request.CreateAppointment) error {
	apt := &domain.Appointment{
		PatientID: p.PatientID,
		DoctorID:  p.DoctorID,
		Date:      p.Date,
		StartHour: p.StartHour,
		EndHour:   p.EndHour,
		Status:    "pending",
	}
	if err := u.aptRepo.Create(apt); err != nil {
		return err
	}

	return nil
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

func (u *AppointmentService) FindAllByPatient(userID string) ([]*domain.Appointment, error) {
	patient, err := u.userRepo.FindPatientProfileByID(userID)
	if err != nil {
		return nil, err
	}

	appointments, err := u.aptRepo.FindAllByPatient(patient.ID)
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
