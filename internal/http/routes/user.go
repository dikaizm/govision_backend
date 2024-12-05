package routes

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/http/middleware"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	router.Handle(
		"/user/profile/patient",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.CreatePatientProfile)),
	).Methods("POST")

	router.Handle(
		"/user/profile/patient",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.ViewPatientProfile)),
	).Methods("GET")

	router.Handle(
		"/user/profile/doctor",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.ViewDoctorProfile)),
	).Methods("GET")

	// ===================== Doctor =====================
	// Doctors for appointment booking

	/*
		@desc Get all doctors
		@route /user/doctor/profile?start_date={start_date}&end_date={end_date}&start_hour={start_hour}&end_hour={end_hour}
		@method GET
	*/
	router.Handle(
		"/user/doctors",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.ViewAll)),
	).Methods("GET")

	/*
		@desc Get doctor profile by user id
		@route /user/doctors/{user_id}
		@method GET
	*/
	router.Handle(
		"/user/doctors/{user_id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.View)),
	).Methods("GET")

	/*
		@desc Get doctor time slots by user id and date
		@route /user/doctors/{user_id}/time-slots/{date}
		@method GET
	*/
	router.Handle(
		"/user/doctors/{user_id}/time-slots/{date}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.GetTimeSlots)),
	).Methods("GET")

	// ===================== End Doctor =====================

	/*
		@desc Get doctor profile by patient
		@route /user/doctor/profile/{id}
		@method GET
	*/
	router.Handle(
		"/user/doctor/profile/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.Profile)),
	).Methods("GET")

	/*
		@desc Create available schedule for doctor
		@route /user/doctor/schedule
		@method POST
		@body []{ "day", "start_hour", "end_hour" }
	*/
	router.Handle(
		"/user/doctor/schedule",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.CreateSchedule)),
	).Methods("POST")

	/*
		@desc Update schedule for doctor
		@route /user/doctor/schedule
		@method PUT
		@body { "id" }
	*/
}
