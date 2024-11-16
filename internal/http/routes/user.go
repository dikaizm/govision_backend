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
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.ViewPatientProfile)),
	).Methods("GET")

	router.Handle(
		"/user/profile/doctor",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.ViewDoctorProfile)),
	).Methods("GET")

	/*
		@desc Get all doctor profile by patient
		@route /user/doctor/profile?start_date={start_date}&end_date={end_date}&start_hour={start_hour}&end_hour={end_hour}
		@method GET
	*/
	router.Handle(
		"/user/doctor/profile",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.ViewAll)),
	).Methods("GET")

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
