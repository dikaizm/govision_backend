package service_intf

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	"github.com/dikaizm/govision_backend/pkg/domain"
)

type AuthService interface {
	Register(p *request.Register) (*response.Register, error)
	Login(p *request.Login) (*response.Login, error)

	RegisterAsDoctor(userID string, p *request.RegisterDoctor) error
	RegisterAsPatient(userID string, p *request.RegisterPatient) error
}

type AppointmentService interface {
	Create(p *request.CreateAppointment) error
	FindAllByDoctor(userID string) ([]*domain.Appointment, error)
	FindAllByPatient(userID string) ([]*domain.Appointment, error)
	UpdateStatus(aptID int64, confirm bool) error
}

type FundusService interface {
	DetectImage(p *request.DetectFundusImage) (res *response.DetectFundusImage, err error)
	ViewFundus(fundusID int64) (*domain.Fundus, error)
	FundusHistory(userID string) ([]*response.FundusHistory, error)
	RequestVerifyFundusByPatient() error
	VerifyFundusByDoctor(fundusID int64, doctorID int64, status string, notes string) error
	DeleteFundus(fundusID int64) error
}

type UserService interface {
	GetProfilePatient(userID string) (*response.GetProfilePatient, error)
	GetProfileDoctor(userID string) (*response.GetProfile, error)
	UpdateProfile() error
}

type DoctorService interface {
	FindAll(filter *request.FilterAppointmentSchedule) ([]*domain.UserDoctor, error)
	GetProfile(profileID int64) (*domain.UserDoctor, error)
	CreateSchedule(userID string, params []*request.CreateDoctorSchedule) error
}
