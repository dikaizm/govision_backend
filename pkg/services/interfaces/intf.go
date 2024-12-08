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

type ArticleService interface {
	Create(p *request.CreateArticle) error
	CreateBulk(p []*request.CreateArticle) error
	FindAll(filter *request.FilterGetArticle) ([]*domain.Article, error)
	FindByID(id string) (*domain.Article, error)
}

type AppointmentService interface {
	Create(p *request.CreateAppointment) (*domain.Appointment, error)
	FindAllByDoctor(userID string) ([]*domain.Appointment, error)
	FindAllByPatient(userID string, filter *request.FilterViewAllAppointment) ([]*domain.Appointment, error)
	UpdateStatus(aptID int64, confirm bool) error
}

type FundusService interface {
	DetectImage(p *request.DetectFundusImage) (res *response.DetectFundusImage, err error)
	ViewFundus(fundusID int64) (*domain.Fundus, error)
	ViewFundusHistory(userID string) ([]*domain.Fundus, error)
	RequestVerifyFundusByPatient(fundusID int64) error
	VerifyFundusByDoctor(fundusID int64, doctorID int64, status string, notes string) error
	DeleteFundus(fundusID int64) error
	GetFundusImage(path string) (string, error)
	ViewVerifiedFundus(userID string) (*response.ViewVerifiedFundus, error)
}

type UserService interface {
	Get(userID string) (*domain.User, error)
	GetProfilePatient(userID string) (*response.GetProfilePatient, error)
	GetProfileDoctor(userID string) (*response.GetProfile, error)
	UpdateProfile() error
}

type DoctorService interface {
	FindAll(filter *request.FilterAppointmentSchedule) ([]*domain.UserDoctor, error)
	GetProfile(userID string) (*domain.UserDoctor, error)
	CreateSchedule(userID string, params []*request.CreateDoctorSchedule) error
	GetTimeSlots(userID string, date string) ([]*domain.DoctorScheduleTimeSlot, error)
}
