package repo_intf

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
)

type ArticleRepository interface {
	Create(article *domain.Article) error
	CreateBulk(articles []*domain.Article) error
	FindAll(filter *request.FilterGetArticle) ([]*domain.Article, error)
	FindByID(id string) (*domain.Article, error)
	Update(article *domain.Article) error
}

type AppointmentRepository interface {
	Create(apt *domain.Appointment) (*domain.Appointment, error)
	FindAllByDoctor(doctorID int64) ([]*domain.Appointment, error)
	FindAllByPatient(patientID int64, filter *request.FilterViewAllAppointment) ([]*domain.Appointment, error)
	UpdateStatus(id int64, status string) error
	Delete(id int64) error
}

type FundusRepository interface {
	Create(fundus *domain.CreateFundus) (*domain.Fundus, error)
	CreateFeedbackByDoctor(fundusID int64, doctorID int64, notes string) error
	FindAllByPatient(patientID int64) (res []*domain.Fundus, err error)
	FindByID(id int64) (*domain.Fundus, error)
	DeleteByID(id int64) error
	DeleteFeedbackByDoctor(id int64, doctorID int64) error
	RequestVerifyStatusByPatient(id int64) error
	UpdateVerifyStatusByDoctor(id int64, doctorID int64, verifyStatus string) error
	FindLastVerifiedByPatient(patientID int64) (*domain.Fundus, error)
}

type UserRepository interface {
	Create(*domain.User) (*string, error)
	FindByID(id string) (*domain.User, error)
	FindPatientProfileByID(id string) (*domain.UserPatient, error)
	FindDoctorProfileByID(id string) (*domain.UserDoctor, error)
	FindByEmail(string) (*domain.User, error)
	GetAllRole() ([]*domain.UserRole, error)

	CreateDoctorProfile(profile *domain.UserDoctor, practices []*domain.DoctorExperience, educations []*domain.DoctorEducation) (*string, error)
	CreatePatientProfile(profile *domain.UserPatient) (*string, error)
}

type DoctorRepository interface {
	CreateProfile(profile *domain.UserDoctor, practices []*domain.DoctorExperience, educations []*domain.DoctorEducation) (*int64, error)
	FindAll(filter *request.FilterAppointmentSchedule) ([]*domain.UserDoctor, error)
	FindProfileByUserID(userID string) (*domain.UserDoctor, error)
	GetPractice(profileID int64) ([]*domain.DoctorExperience, error)
	GetSchedule(profileID int64) ([]*domain.DoctorSchedule, error)
	CreateSchedule(schedules []*domain.DoctorSchedule) error
	FindTimeSlotsByProfileIDAndDate(profileID int64, date string) ([]*domain.DoctorScheduleTimeSlot, error)
	UpdateTimeSlotToBooked(timeSlotID int64) error
}
