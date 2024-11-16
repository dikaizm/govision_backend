package routes

import (
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router, controller route_intf.Controllers) {
	router.HandleFunc("/auth/register/complete", controller.Auth.Register).Methods("POST")
	router.HandleFunc("/auth/login", controller.Auth.Login).Methods("POST")
}
