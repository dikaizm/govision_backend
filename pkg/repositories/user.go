package repositories

import (
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type DbUserRepository struct {
	DB *gorm.DB
}

func NewDbUserRepository(db *gorm.DB) repo_intf.UserRepository {
	return &DbUserRepository{
		DB: db,
	}
}

func (r *DbUserRepository) Create(user *domain.User) (*string, error) {
	var userCode string = ""
	if err := r.DB.Create(user).Error; err != nil {
		return &userCode, err
	}
	return &user.ID, nil
}

func (r *DbUserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Preload("Role").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *DbUserRepository) FindPatientProfileByID(id string) (*domain.UserPatient, error) {
	var profile domain.UserPatient
	if err := r.DB.Where("user_id = ?", id).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *DbUserRepository) FindDoctorProfileByID(id string) (*domain.UserDoctor, error) {
	var profile domain.UserDoctor
	if err := r.DB.Where("user_id = ?", id).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *DbUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *DbUserRepository) GetAllRole() ([]*domain.UserRole, error) {
	var roles []*domain.UserRole
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *DbUserRepository) CreateDoctorProfile(profile *domain.UserDoctor, practices []*domain.DoctorExperience, educations []*domain.DoctorEducation) (*string, error) {
	tx := r.DB.Begin()
	if err := tx.Create(profile).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, practice := range practices {
		practice.ProfileID = profile.ID
		if err := tx.Create(practice).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Update("work_years", profile.WorkYears).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, education := range educations {
		education.ProfileID = profile.ID
		if err := tx.Create(education).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &profile.UserID, nil
}

func (r *DbUserRepository) CreatePatientProfile(profile *domain.UserPatient) (*string, error) {
	if err := r.DB.Create(profile).Error; err != nil {
		return nil, err
	}
	return &profile.UserID, nil
}
