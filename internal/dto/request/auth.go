package request

import (
	"github.com/dikaizm/govision_backend/pkg/helpers/dtype"
)

type (
	Register struct {
		Name            string `json:"name" validate:"required,min=3,max=100"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8,max=100"`
		ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=100"`

		Phone     string     `json:"phone" validate:"required,e164"`
		BirthDate dtype.Date `json:"birth_date" validate:"required"`
		Gender    string     `json:"gender" validate:"required,oneof=male female"`

		Role string `json:"role" validate:"required,oneof=doctor patient"`

		Village       string `json:"village"`
		Subdistrict   string `json:"subdistrict"`
		City          string `json:"city" validate:"required"`
		Province      string `json:"province" validate:"required"`
		AddressDetail string `json:"address_detail" validate:"required"`
	}
)

type (
	DoctorEducation struct {
		University string `json:"university" validate:"required,min=3,max=100"`
		Major      string `json:"major" validate:"required,min=3,max=100"`
		StartYear  int    `json:"start_year" validate:"required"`
		EndYear    int    `json:"end_year" validate:"required"`
	}

	DoctorPractice struct {
		InstitutionName string     `json:"institution_name" validate:"required,min=3,max=100"`
		City            string     `json:"city" validate:"required,min=3,max=100"`
		Province        string     `json:"province" validate:"required,min=3,max=100"`
		AddressDetail   string     `json:"address_detail" validate:"required,min=3,max=255"`
		StartDate       dtype.Date `json:"start_date" validate:"required"`
		EndDate         dtype.Date `json:"end_date" validate:"required"`
	}

	RegisterDoctor struct {
		Specialization string            `json:"specialization" validate:"required,min=3,max=100"`
		StrNo          string            `json:"str_no" validate:"required,min=3,max=100"`
		BioDesc        string            `json:"bio_desc" validate:"required,min=3,max=255"`
		WorkYears      int               `json:"work_years" validate:"required"`
		Educations     []DoctorEducation `json:"educations" validate:"required"`
		Practices      []DoctorPractice  `json:"practices" validate:"required"`
	}
)

type (
	RegisterPatient struct {
		DiabetesHistory bool       `json:"diabetes_history" validate:"required"`
		DiabetesType    string     `json:"diabetes_type" validate:"required"`
		DiagnosisDate   dtype.Date `json:"diagnosis_date" validate:"required"`
	}
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
