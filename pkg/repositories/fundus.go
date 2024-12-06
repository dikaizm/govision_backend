package repositories

import (
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type DbFundusRepository struct {
	DB *gorm.DB
}

func NewDbFundusRepository(db *gorm.DB) repo_intf.FundusRepository {
	return &DbFundusRepository{DB: db}
}

func (r *DbFundusRepository) Create(fundus *domain.CreateFundus) (*domain.Fundus, error) {
	if err := r.DB.Table("fundus").Create(fundus).Error; err != nil {
		return nil, err
	}

	newFundus, err := r.FindByID(fundus.ID)
	if err != nil {
		return nil, err
	}

	return newFundus, nil
}

func (r *DbFundusRepository) CreateFeedbackByDoctor(fundusID int64, doctorID int64, notes string) error {
	feedback := &domain.FundusFeedback{
		FundusID: fundusID,
		DoctorID: doctorID,
		Notes:    notes,
	}

	if err := r.DB.Create(feedback).Error; err != nil {
		return err
	}
	return nil
}

func (r *DbFundusRepository) FindAllByPatient(patientID int64) (res []*domain.Fundus, err error) {
	if err = r.DB.Where("patient_id = ?", patientID).
		Order("created_at desc, updated_at desc").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *DbFundusRepository) FindByID(id int64) (*domain.Fundus, error) {
	var fundus *domain.Fundus
	err := r.DB.
		Preload("Feedbacks").
		Preload("Feedbacks.Doctor").
		Preload("Feedbacks.Doctor.User").
		First(&fundus, id).Error
	if err != nil {
		return nil, err
	}
	return fundus, nil
}

func (r *DbFundusRepository) DeleteByID(id int64) error {
	if err := r.DB.Delete(&domain.Fundus{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *DbFundusRepository) DeleteFeedbackByDoctor(id int64, doctorID int64) error {
	err := r.DB.Where("id = ? AND doctor_id = ?", id, doctorID).Delete(&domain.FundusFeedback{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DbFundusRepository) UpdateVerifyStatusByDoctor(id int64, doctorID int64, verifyStatus string) error {
	err := r.DB.Model(&domain.Fundus{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verified_by":   doctorID,
			"verify_status": verifyStatus,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DbFundusRepository) FindLastVerifiedByPatient(patientID int64) (*domain.Fundus, error) {
	var fundus *domain.Fundus

	err := r.DB.
		Where("patient_id = ?", patientID).
		Where("verify_status = ?", "verified").
		Order("updated_at desc").
		First(&fundus).Error
	if err != nil {
		return nil, err
	}

	return fundus, nil
}
