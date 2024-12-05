package services

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type DoctorService struct {
	doctorRepo repo_intf.DoctorRepository
	userRepo   repo_intf.UserRepository
}

func NewDoctorService(doctorRepo repo_intf.DoctorRepository, userRepo repo_intf.UserRepository) service_intf.DoctorService {
	return &DoctorService{
		doctorRepo: doctorRepo,
		userRepo:   userRepo,
	}
}

func (u *DoctorService) FindAll(filter *request.FilterAppointmentSchedule) ([]*domain.UserDoctor, error) {
	doctors, err := u.doctorRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}

func (u *DoctorService) GetProfile(userID string) (*domain.UserDoctor, error) {
	doctor, err := u.doctorRepo.FindProfileByUserID(userID)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (u *DoctorService) CreateSchedule(userID string, params []*request.CreateDoctorSchedule) error {
	var schedules []*domain.DoctorSchedule

	_, err := u.userRepo.FindDoctorProfileByID(userID)
	if err != nil {
		return err
	}

	err = u.doctorRepo.CreateSchedule(schedules)
	if err != nil {
		return err
	}

	return nil
}

func (u *DoctorService) GetTimeSlots(userID string, date string) ([]*domain.DoctorScheduleTimeSlot, error) {
	profile, err := u.userRepo.FindDoctorProfileByID(userID)
	if err != nil {
		return nil, err
	}

	timeSlots, err := u.doctorRepo.FindTimeSlotsByProfileIDAndDate(profile.ID, date)
	if err != nil {
		return nil, err
	}

	return timeSlots, nil
}
