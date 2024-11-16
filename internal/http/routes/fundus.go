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
		@desc Get a fundus by user
		@route /fundus/{id}
		@method GET
	*/
	router.Handle(
		"/fundus/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.ViewFundus)),
	).Methods("GET")

	/*
		@route /fundus/verify/{id}
		@method POST
		@body { "doctor_id", "status", "[]feedbacks" }
	*/
	router.Handle(
		"/fundus/verify/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.VerifyFundusByDoctor)),
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
