package domain

import "time"

const (
	FundusDiseaseNoDR        = "no_dr"
	FundusDiseaseMild        = "mild_dr"
	FundusDiseaseModerate    = "moderate_dr"
	FundusDiseaseSevere      = "severe_dr"
	FundusDiseaseProliferate = "proliferate_dr"

	FundusDiseaseNotDetected = "not_detected"
)

const (
	FundusVerifyStatusPending  = "pending"
	FundusVerifyStatusOnReview = "on_review"
	FundusVerifyStatusVerified = "verified"
	FundusVerifyStatusRejected = "rejected"
)

type Fundus struct {
	ID               int64       `gorm:"primaryKey"`
	PatientID        int64       `gorm:"not null"`
	Patient          UserPatient `gorm:"foreignKey:PatientID"`
	ImgURL           string      `gorm:"size:255"`
	VerifiedBy       int64       `gorm:"null"`
	Verifier         UserDoctor  `gorm:"foreignKey:VerifiedBy;null"`
	VerifyStatus     string      `gorm:"size:255"`
	PredictedDisease string      `gorm:"size:20;check:predicted_disease IN ('no_dr','mild_dr','moderate_dr','severe_dr','proliferative_dr', 'not_detected')"`
	CreatedAt        time.Time   `gorm:"autoCreateTime"`
	UpdatedAt        time.Time   `gorm:"autoUpdateTime"`

	Feedbacks []FundusFeedback `gorm:"foreignKey:FundusID;references:ID"`
}

type CreateFundus struct {
	ID               int64     `gorm:"primaryKey"`
	PatientID        int64     `gorm:"not null"`
	ImgURL           string    `gorm:"size:255"`
	VerifyStatus     string    `gorm:"size:255"`
	PredictedDisease string    `gorm:"size:20;check:predicted_disease IN ('no_dr','mild_dr','moderate_dr','severe_dr','proliferative_dr', 'not_detected')"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

type FundusFeedback struct {
	ID        int64      `gorm:"primaryKey"`
	FundusID  int64      `gorm:"not null"`
	Fundus    Fundus     `gorm:"foreignKey:FundusID"`
	DoctorID  int64      `gorm:"not null"`
	Doctor    UserDoctor `gorm:"foreignKey:DoctorID"`
	Notes     string     `gorm:"size:255"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
