package routes

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/http/middleware"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func ArticleRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	router.HandleFunc("/articles", controller.Article.ViewAll).Methods("GET")
	router.HandleFunc("/articles/{id}", controller.Article.View).Methods("GET")

	router.Handle(
		"/articles",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Article.Create)),
	).Methods("POST")
	router.Handle(
		"/articles/bulk",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Article.CreateBulk)),
	)
}
