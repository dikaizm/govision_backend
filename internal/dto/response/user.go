package response

import (
	"time"

	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
)

type GetProfile struct {
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Address   string    `json:"address"`
}

type GetProfilePatient struct {
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Role          string     `json:"role"`
	BirthDate     dtype.Date `json:"birth_date"`
	Gender        string     `json:"gender"`
	City          string     `json:"city"`
	Province      string     `json:"province"`
	AddressDetail string     `json:"address_detail"`

	DiabetesHistory bool       `json:"diabetes_history"`
	DiabetesType    string     `json:"diabetes_type"`
	DiagnosisDate   dtype.Date `json:"diagnosis_date"`
}

type (
	CurrentUser struct {
		ID   int64  `json:"id"`
		Role string `json:"role"`
	}
)

type GetUser struct {
	UserID           string `json:"user_id"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	Email            string `json:"email"`
	Photo            string `json:"photo"`
	CompletedProfile bool   `json:"completed_profile"`
}
