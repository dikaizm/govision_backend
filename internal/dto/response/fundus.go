package response

import (
	"time"

	"github.com/dikaizm/govision_backend/pkg/domain"
)

type (
	Fundus struct {
		ID        int64                   `json:"id"`
		ImageBlob string                  `json:"image_blob"`
		Verified  bool                    `json:"verified"`
		Status    string                  `json:"status"`
		Condition string                  `json:"condition"`
		CreatedAt string                  `json:"created_at"`
		UpdatedAt string                  `json:"updated_at,omitempty"`
		Feedbacks []domain.FundusFeedback `json:"feedbacks,omitempty"`
	}

	FundusID struct {
		ID int64 `json:"fundus_id"`
	}
)

type (
	FundusVerifier struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Specialization string `json:"specialization"`
	}

	FundusFeedback struct {
		ID           int64     `json:"id"`
		DoctorUserID string    `json:"doctor_user_id"`
		DoctorName   string    `json:"doctor_name"`
		Notes        string    `json:"notes"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	}

	DetectFundusImage struct {
		ID               int64            `json:"id"`
		ImageBase64      string           `json:"image_base64"`
		VerifyStatus     string           `json:"verify_status"`
		Verifier         FundusVerifier   `json:"verifier"`
		PredictedDisease string           `json:"predicted_disease"`
		CreatedAt        time.Time        `json:"created_at"`
		UpdatedAt        time.Time        `json:"updated_at,omitempty"`
		Feedbacks        []FundusFeedback `json:"feedbacks,omitempty"`
	}

	FundusHistory struct {
		ID               int64     `json:"id"`
		ImageBase64      string    `json:"image_base64"`
		VerifyStatus     string    `json:"verify_status"`
		PredictedDisease string    `json:"predicted_disease"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at,omitempty"`
	}

	ViewFundusHistory struct {
		ID               int64            `json:"id"`
		ImageUrl         string           `json:"image_url"`
		VerifyStatus     string           `json:"verify_status"`
		PredictedDisease string           `json:"predicted_disease"`
		CreatedAt        time.Time        `json:"created_at"`
		UpdatedAt        time.Time        `json:"updated_at,omitempty"`
		Feedbacks        []FundusFeedback `json:"feedbacks"`
	}
)

type (
	ViewVerifiedFundus struct {
		ID               int64  `json:"id"`
		ImageUrl         string `json:"image_url"`
		VerifyStatus     string `json:"verify_status"`
		PredictedDisease string `json:"predicted_disease"`

		DiabetesType string `json:"diabetes_type"`

		RecommendedExamination string `json:"recommended_examination"`
		RecommendedNotes       string `json:"recommended_notes"`
	}
)
