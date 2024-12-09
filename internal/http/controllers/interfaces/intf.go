package controller_intf

import (
	"net/http"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type ArticleController interface {
	Create(w http.ResponseWriter, r *http.Request)
	CreateBulk(w http.ResponseWriter, r *http.Request)
	ViewAll(w http.ResponseWriter, r *http.Request)
	View(w http.ResponseWriter, r *http.Request)
}

type AppointmentController interface {
	Create(w http.ResponseWriter, r *http.Request)
	ViewAll(w http.ResponseWriter, r *http.Request)
	Confirm(w http.ResponseWriter, r *http.Request)
}

type FundusController interface {
	DetectFundusImage(w http.ResponseWriter, r *http.Request)
	ViewFundusHistory(w http.ResponseWriter, r *http.Request)
	ViewFundus(w http.ResponseWriter, r *http.Request)
	DeleteFundus(w http.ResponseWriter, r *http.Request)

	RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request)
	UpdateVerifyFundusByDoctor(w http.ResponseWriter, r *http.Request)

	ViewFundusImage(w http.ResponseWriter, r *http.Request)
	ViewVerifiedFundus(w http.ResponseWriter, r *http.Request)
}

type UserController interface {
	View(w http.ResponseWriter, r *http.Request)

	ViewPatientProfile(w http.ResponseWriter, r *http.Request)
	ViewDoctorProfile(w http.ResponseWriter, r *http.Request)

	CreatePatientProfile(w http.ResponseWriter, r *http.Request)
	CreateDoctorProfile(w http.ResponseWriter, r *http.Request)
}

type DoctorController interface {
	ViewAll(w http.ResponseWriter, r *http.Request)
	View(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	CreateSchedule(w http.ResponseWriter, r *http.Request)
	GetTimeSlots(w http.ResponseWriter, r *http.Request)
}
