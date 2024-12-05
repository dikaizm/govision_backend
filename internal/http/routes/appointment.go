package routes

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/http/middleware"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func AppointmentRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	/*
		@route /appointment
		@method GET
	*/
	router.Handle(
		"/appointments",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Appointment.ViewAll)),
	).Methods("GET")

	/*
		@desc Create appointment
		@route /appointment
		@method POST
	*/
	router.Handle(
		"/appointments",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Appointment.Create)),
	).Methods("POST")

	/*
		@desc Confirm appointment by doctor
		@router /appointment/confirm/{apt_id}
		@method POST
		@body { "confirm" }
	*/
	router.Handle(
		"/appointments/confirm/{apt_id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Appointment.Confirm)),
	).Methods("POST")
}
