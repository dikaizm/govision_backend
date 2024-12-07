package driver_db

import (
	"fmt"
	"log"

	"github.com/dikaizm/govision_backend/pkg/config"
	"github.com/dikaizm/govision_backend/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SslMode  string
}

func NewConnection(env *config.Env) (*gorm.DB, error) {
	config := ConfigDB{
		Host:     env.DbHost,
		Port:     env.DbPort,
		User:     env.DbUser,
		Name:     env.DbName,
		Password: env.DbPassword,
		SslMode:  env.DbSslMode,
	}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.Name, config.Port, config.SslMode)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	log.Println("Connected to postgres database")
	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
		&domain.UserRole{},
		&domain.UserDoctor{},
		&domain.UserPatient{},
		&domain.DoctorExperience{},
		&domain.DoctorEducation{},
		&domain.DoctorSchedule{},
		&domain.DoctorScheduleTimeSlot{},
		&domain.Appointment{},
		&domain.Fundus{},
		&domain.FundusFeedback{},
		&domain.Article{},
	)
}

func SeedRole(db *gorm.DB) {
	roles := []domain.UserRole{
		{RoleName: "admin"},
		{RoleName: "patient"},
		{RoleName: "doctor"},
	}

	for _, role := range roles {
		// Check if the role already exists in the database
		var existingRole domain.UserRole
		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Role not found, create it
				db.Create(&role)
			} else {
				// Handle other errors (if any)
				fmt.Println("Error checking role:", err)
			}
		}
	}
}
