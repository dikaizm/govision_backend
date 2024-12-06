package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dikaizm/govision_backend/internal/http/controllers"
	"github.com/dikaizm/govision_backend/internal/http/routes"
	route_intf "github.com/dikaizm/govision_backend/internal/http/routes/interfaces"
	"github.com/dikaizm/govision_backend/pkg/config"
	driver_db "github.com/dikaizm/govision_backend/pkg/driver/db"
	"github.com/dikaizm/govision_backend/pkg/repositories"
	"github.com/dikaizm/govision_backend/pkg/services"
)

func main() {
	log.Print("Starting server...")

	env := config.LoadEnv()

	db, err := driver_db.NewConnection(env)
	if err != nil {
		log.Println(err)
	}

	driver_db.AutoMigrate(db)

	userRepo := repositories.NewDbUserRepository(db)
	doctorRepo := repositories.NewDbDoctorRepository(db)
	fundusRepo := repositories.NewDbFundusRepository(db)
	appointmentRepo := repositories.NewDbAppointmentRepository(db)
	articleRepo := repositories.NewDbArticleRepository(db)

	authService := services.NewAuthService(env.SecretKey, userRepo)
	articleService := services.NewArticleService(articleRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo, userRepo, doctorRepo)
	doctorService := services.NewDoctorService(doctorRepo, userRepo)
	fundusService := services.NewFundusService(env.MlApi, env.MlApiKey, fundusRepo, userRepo)
	userService := services.NewUserService(userRepo)

	authController := controllers.NewAuthController(authService)
	articleController := controllers.NewArticleController(articleService)
	appointmentController := controllers.NewAppointmentController(appointmentService)
	doctorController := controllers.NewDoctorController(doctorService)
	fundusController := controllers.NewFundusController(fundusService)
	userController := controllers.NewUserController(userService, authService)

	router := routes.SetupRouter(env.SecretKey, route_intf.Controllers{
		Auth:        authController,
		Article:     articleController,
		Appointment: appointmentController,
		Doctor:      doctorController,
		Fundus:      fundusController,
		User:        userController,
	})

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
	http.HandleFunc("/hello-world", helloHandler)

	port := fmt.Sprintf(":%s", env.AppPort)

	log.Println("Server started at: ", env.AppPort)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	}
}
