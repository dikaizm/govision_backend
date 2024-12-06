package routes

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/http/middleware"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func FundusRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	router.Handle(
		"/fundus/detect",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.DetectFundusImage)),
	).Methods("POST")

	/*
		@desc Get all fundus
		@route /fundus
		@method GET
	*/
	router.Handle(
		"/fundus",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.ViewFundusHistory)),
	).Methods("GET")

	/*
		@desc Get a fundus by user
		@route /fundus/{id}
		@method GET
	*/
	router.Handle(
		"/fundus/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.ViewFundus)),
	).Methods("GET")

	/*
		@desc Get fundus image by path for private access by patient and doctor
		@route /fundus/image/{path}
		@method GET
		@type public
	*/
	router.Handle(
		"/fundus/image/{path}",
		http.HandlerFunc(controller.Fundus.ViewFundusImage),
	).Methods("GET")

	/*
		@desc Get last verified fundus
		@route /fundus/home/verified
		@method GET
	*/
	router.Handle(
		"/fundus/home/verified",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.ViewVerifiedFundus)),
	).Methods("GET")

	router.Handle(
		"/fundus/{id}/request-verify",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.RequestVerifyFundusByPatient)),
	).Methods("POST")

	/*
		@route /fundus/set-verify/{id}
		@method POST
		@body { "doctor_id", "status", "[]feedbacks" }
	*/
	router.Handle(
		"/fundus/{id}/update-verify",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.UpdateVerifyFundusByDoctor)),
	).Methods("POST")

	/*
		@route /fundus/{id}
		@method DELETE
	*/
	router.Handle(
		"/fundus/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.DeleteFundus)),
	).Methods("DELETE")
}
