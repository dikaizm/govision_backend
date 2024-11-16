package services

import (
	"github.com/dikaizm/govision_backend/internal/dto/response"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type UserService struct {
	userRepo repo_intf.UserRepository
}

func NewUserService(userRepo repo_intf.UserRepository) service_intf.UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetProfilePatient(userID string) (*response.GetProfilePatient, error) {
	var profile response.GetProfilePatient

	patient, err := s.userRepo.FindPatientProfileByID(userID)
	if err != nil {
		return nil, err
	}

	profile = response.GetProfilePatient{
		Name:          patient.User.Name,
		Email:         patient.User.Email,
		Phone:         patient.User.Phone,
		Role:          patient.User.Role.RoleName,
		BirthDate:     patient.User.BirthDate,
		Gender:        patient.User.Gender,
		City:          patient.User.City,
		Province:      patient.User.Province,
		AddressDetail: patient.User.AddressDetail,

		DiabetesHistory: patient.DiabetesHistory,
		DiabetesType:    patient.DiabetesType,
		DiagnosisDate:   patient.DiagnosisDate,
	}

	return &profile, nil
}

func (s *UserService) GetProfileDoctor(userID string) (*response.GetProfile, error) {
	return nil, nil
}

func (s *UserService) UpdateProfile() error {
	return nil
}
