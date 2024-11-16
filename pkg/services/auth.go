package services

import (
	"errors"
	"log"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	"github.com/dikaizm/govision_backend/pkg/domain"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type AuthService struct {
	secretKey string
	userRepo  repo_intf.UserRepository
}

func NewAuthService(secretKey string, userRepo repo_intf.UserRepository) service_intf.AuthService {
	return &AuthService{
		secretKey: secretKey,
		userRepo:  userRepo,
	}
}

func (u *AuthService) Register(p *request.Register) (*response.Register, error) {
	// Check if password and confirm password match
	if p.Password != p.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(u.secretKey, p.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &domain.User{
		ID:            helpers.GenerateUserID(),
		Name:          p.Name,
		Phone:         p.Phone,
		Email:         p.Email,
		Password:      hashedPassword,
		RoleID:        p.RoleID,
		BirthDate:     p.BirthDate,
		Gender:        p.Gender,
		City:          p.City,
		Province:      p.Province,
		AddressDetail: p.AddressDetail,
	}

	_, err = u.userRepo.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// Generate JWT
	paramsJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60 * 24 * 30, // valid for 30 days
		SecretKey:       u.secretKey,
		UserID:          user.ID,
		UserRole:        p.Role,
	}

	resultJWT, err := helpers.GenerateJWT(&paramsJWT)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &response.Register{
		Name:        p.Name,
		Email:       p.Email,
		Role:        p.Role,
		AccessToken: resultJWT.Token,
	}, nil
}

func (u *AuthService) Login(p *request.Login) (*response.Login, error) {
	// Find user by email
	user, err := u.userRepo.FindByEmail(p.Email)
	if err != nil {
		log.Printf("Error finding user by email: %v", err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check password
	err = helpers.CheckPasswordHash(u.secretKey, p.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate JWT
	paramsJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60 * 24 * 30,
		SecretKey:       u.secretKey,
		UserID:          user.ID,
		UserRole:        user.Role.RoleName,
	}

	resultJWT, err := helpers.GenerateJWT(&paramsJWT)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &response.Login{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role.RoleName,
		AccessToken: resultJWT.Token,
	}, nil
}

func (u *AuthService) RegisterAsDoctor(userID string, p *request.RegisterDoctor) error {
	profile := &domain.UserDoctor{
		UserID:         userID,
		Specialization: p.Specialization,
		StrNo:          p.StrNo,
		BioDesc:        p.BioDesc,
		WorkYears:      0,
		Rating:         0,
	}

	practices := []*domain.DoctorPractice{}
	for _, pr := range p.Practices {
		workYears := helpers.GetWorkYears(pr.StartDate, pr.EndDate)

		practice := &domain.DoctorPractice{
			City:       pr.City,
			Province:   pr.Province,
			OfficeName: pr.OfficeName,
			StartDate:  pr.StartDate,
			EndDate:    pr.EndDate,
		}
		practices = append(practices, practice)
		profile.WorkYears += workYears
	}

	educations := []*domain.DoctorEducation{}
	for _, ed := range p.Educations {
		education := &domain.DoctorEducation{
			Major:      ed.Major,
			University: ed.University,
			StartYear:  ed.StartYear,
			EndYear:    ed.EndYear,
		}
		educations = append(educations, education)
	}

	_, err := u.userRepo.CreateDoctorProfile(profile, practices, educations)
	if err != nil {
		return err
	}

	return nil
}

func (u *AuthService) RegisterAsPatient(userID string, p *request.RegisterPatient) error {
	profile := &domain.UserPatient{
		UserID:          userID,
		DiabetesHistory: p.DiabetesHistory,
		DiabetesType:    p.DiabetesType,
		DiagnosisDate:   p.DiagnosisDate,
	}

	_, err := u.userRepo.CreatePatientProfile(profile)
	if err != nil {
		return err
	}

	return nil
}
