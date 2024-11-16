package services

import (
	"errors"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	"github.com/dikaizm/govision_backend/pkg/helpers"
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
	daysOfWeek, err := helpers.GetDaysOfWeek(filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, errors.New("failed to get days of week")
	}
	filter.DaysInt = daysOfWeek

	doctors, err := u.doctorRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}

func (u *DoctorService) GetProfile(doctorID int64) (*domain.UserDoctor, error) {
	doctor, err := u.doctorRepo.GetProfileByID(doctorID)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (u *DoctorService) CreateSchedule(userID string, params []*request.CreateDoctorSchedule) error {
	var schedules []*domain.DoctorSchedule

	doctor, err := u.userRepo.FindDoctorProfileByID(userID)
	if err != nil {
		return err
	}

	for _, p := range params {
		schedule := &domain.DoctorSchedule{
			ProfileID: doctor.ID,
			DayOfWeek: p.DayOfWeek,
			StartHour: p.StartHour,
			EndHour:   p.EndHour,
		}
		schedules = append(schedules, schedule)
	}

	err = u.doctorRepo.CreateSchedule(schedules)
	if err != nil {
		return err
	}

	return nil
}
