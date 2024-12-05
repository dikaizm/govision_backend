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

func (r *DbDoctorRepository) CreateProfile(profile *domain.UserDoctor, practices []*domain.DoctorExperience, educations []*domain.DoctorEducation) (*int64, error) {
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
	tx := r.DB.Model(&domain.UserDoctor{}).Where("is_verified = ?", true)

	if filter != nil {
		if len(filter.Days) > 0 || filter.StartHour != "" || filter.EndHour != "" {
			tx = tx.Joins("JOIN doctor_schedules ON doctor_schedules.profile_id = user_doctors.id")
		}

		if len(filter.Days) > 0 {
			tx = tx.Where("doctor_schedules.day_of_week IN ?", filter.Days)
		}
		if filter.StartHour != "" {
			tx = tx.Where("doctor_schedules.start_hour >= ?", filter.StartHour)
		}
		if filter.EndHour != "" {
			tx = tx.Where("doctor_schedules.end_hour <= ?", filter.EndHour)
		}
	}

	tx = tx.Preload("Schedules").Preload("User")

	if err := tx.Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *DbDoctorRepository) FindProfileByUserID(userID string) (*domain.UserDoctor, error) {
	var profile domain.UserDoctor
	if err := r.DB.Where("user_id = ?", userID).
		Preload("User").
		Preload("Experiences").
		Preload("Educations").
		Preload("Schedules").
		Preload("Schedules.TimeSlots").
		First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *DbDoctorRepository) GetPractice(profileID int64) ([]*domain.DoctorExperience, error) {
	var practices []*domain.DoctorExperience
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

func (r *DbDoctorRepository) FindTimeSlotsByProfileIDAndDate(profileID int64, date string) ([]*domain.DoctorScheduleTimeSlot, error) {
	var schedule domain.DoctorSchedule
	if err := r.DB.Where("profile_id = ? AND date = ?", profileID, date).First(&schedule).Error; err != nil {
		return nil, err
	}

	if schedule.ID == 0 {
		return nil, nil
	}

	var timeSlots []*domain.DoctorScheduleTimeSlot
	if err := r.DB.Where("schedule_id = ?", schedule.ID).Order("start_time ASC").Find(&timeSlots).Error; err != nil {
		return nil, err
	}

	return timeSlots, nil
}

func (r *DbDoctorRepository) UpdateTimeSlotToBooked(timeSlotID int64) error {
	if err := r.DB.Model(&domain.DoctorScheduleTimeSlot{}).
		Where("id = ?", timeSlotID).
		Update("is_booked", true).
		Error; err != nil {
		return err
	}
	return nil
}
