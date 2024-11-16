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
		&domain.DoctorPractice{},
		&domain.DoctorEducation{},
		&domain.DoctorSchedule{},
		&domain.Appointment{},
		&domain.Fundus{},
		&domain.FundusFeedback{},
	)
}
