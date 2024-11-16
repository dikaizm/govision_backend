package repositories

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type DbDoctorRepository struct {
	DB *gorm.DB
}

func NewDbDoctorRepository(db *gorm.DB) repo_intf.DoctorRepository {
	return &DbDoctorRepository{DB: db}
}

func (r *DbDoctorRepository) CreateProfile(profile *domain.UserDoctor, practices []*domain.DoctorPractice, educations []*domain.DoctorEducation) (*int64, error) {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Create doctor profile
		if err := tx.Create(profile).Error; err != nil {
			return err
		}

		// Create doctor practices
		for _, practice := range practices {
			practice.ProfileID = profile.ID
		}
		if err := tx.Create(&practices).Error; err != nil {
			return err
		}

		// Create doctor educations
		for _, education := range educations {
			education.ProfileID = profile.ID
		}
		if err := tx.Create(&educations).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &profile.ID, nil
}

func (r *DbDoctorRepository) FindAll(filter *request.FilterAppointmentSchedule) ([]*domain.UserDoctor, error) {
	var profiles []*domain.UserDoctor
	tx := r.DB.Model(&domain.UserDoctor{}).Preload("Schedules")

	if filter != nil {
		if len(filter.DaysInt) > 0 {
			tx = tx.Joins("JOIN doctor_schedules ON doctor_schedules.doctor_id = doctor_profiles.id").
				Where("doctor_schedules.day_of_week IN ?", filter.DaysInt)
		}
		if filter.StartHour != "" {
			tx = tx.Where("doctor_schedules.start_hour >= ?", filter.StartHour)
		}
		if filter.EndHour != "" {
			tx = tx.Where("doctor_schedules.end_hour <= ?", filter.EndHour)
		}
	}

	if err := tx.Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *DbDoctorRepository) FindProfileByUserID(userID int64) (*int64, error) {
	var profile domain.UserDoctor
	if err := r.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile.ID, nil
}

func (r *DbDoctorRepository) GetProfileByID(profileID int64) (*domain.UserDoctor, error) {
	var profile domain.UserDoctor
	if err := r.DB.Preload("Practices").Preload("Schedules").
		Where("id = ?", profileID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *DbDoctorRepository) GetPractice(profileID int64) ([]*domain.DoctorPractice, error) {
	var practices []*domain.DoctorPractice
	if err := r.DB.Where("doctor_id = ?", profileID).Find(&practices).Error; err != nil {
		return nil, err
	}
	return practices, nil
}

func (r *DbDoctorRepository) GetSchedule(profileID int64) ([]*domain.DoctorSchedule, error) {
	var schedules []*domain.DoctorSchedule
	if err := r.DB.Where("doctor_id = ?", profileID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *DbDoctorRepository) CreateSchedule(schedules []*domain.DoctorSchedule) error {
	if err := r.DB.Create(&schedules).Error; err != nil {
		return err
	}
	return nil
}
