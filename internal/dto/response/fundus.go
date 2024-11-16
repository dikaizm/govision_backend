package response

import "github.com/dikaizm/govision_backend/pkg/domain"

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
		ID        int64  `json:"id"`
		Notes     string `json:"notes"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}

	DetectFundusImage struct {
		ID               int64            `json:"id"`
		ImageBase64      string           `json:"image_base64"`
		VerifyStatus     string           `json:"verify_status"`
		Verifier         FundusVerifier   `json:"verifier"`
		PredictedDisease string           `json:"predicted_disease"`
		CreatedAt        string           `json:"created_at"`
		UpdatedAt        string           `json:"updated_at,omitempty"`
		Feedbacks        []FundusFeedback `json:"feedbacks,omitempty"`
	}

	FundusHistory struct {
		ID               int64  `json:"id"`
		ImageBase64      string `json:"image_base64"`
		VerifyStatus     string `json:"verify_status"`
		PredictedDisease string `json:"predicted_disease"`
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"updated_at,omitempty"`
	}
)
