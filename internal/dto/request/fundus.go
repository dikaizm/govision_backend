package request

type (
	DetectFundusImage struct {
		UserID      string `json:"user_id"`
		FundusImage string `json:"fundus_image"`
	}

	VerifyFundus struct {
		DoctorID int64  `json:"doctor_id"`
		Status   string `json:"status"`
		Notes    string `json:"notes"`
	}
)
