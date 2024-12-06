package routes

import (
	"github.com/dikaizm/govision_backend/internal/http/middleware"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func SetupRouter(secretKey string, c route_intf.Controllers) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)

	groupRouter := router.PathPrefix("/api").Subrouter()

	// Auth routes
	AuthRoutes(groupRouter, c)
	ArticleRoutes(groupRouter, c, secretKey)
	AppointmentRoutes(groupRouter, c, secretKey)
	FundusRoutes(groupRouter, c, secretKey)
	UserRoutes(groupRouter, c, secretKey)

	return router
}
