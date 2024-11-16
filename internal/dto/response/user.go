package response

import (
	"time"
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
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	BirthDate     time.Time `json:"birth_date"`
	Gender        string    `json:"gender"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	AddressDetail string    `json:"address_detail"`

	DiabetesHistory bool      `json:"diabetes_history"`
	DiabetesType    string    `json:"diabetes_type"`
	DiagnosisDate   time.Time `json:"diagnosis_date"`
}

type (
	CurrentUser struct {
		ID   int64  `json:"id"`
		Role string `json:"role"`
	}
)
