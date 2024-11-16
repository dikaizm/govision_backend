package domain

import (
	"time"
)

const (
	GenderMale   = "male"
	GenderFemale = "female"
)

const (
	UserRolePatient = "patient"
	UserRoleDoctor  = "doctor"
	UserRoleAdmin   = "admin"
)

type User struct {
	ID            string    `gorm:"primaryKey"`
	Name          string    `gorm:"not null;size:100"`
	Phone         string    `gorm:"not null;size:50;unique"`
	Email         string    `gorm:"not null;size:255;unique"`
	Password      string    `gorm:"not null;size:255"`
	RoleID        int       `gorm:"not null"`
	Role          UserRole  `gorm:"foreignKey:RoleID"`
	BirthDate     time.Time `gorm:"not null"`
	Gender        string    `gorm:"size:6;check:gender IN ('male','female')"`
	Village       string    `gorm:"not null;size:100"`
	Subdistrict   string    `gorm:"not null;size:100"`
	City          string    `gorm:"not null;size:100"`
	Province      string    `gorm:"not null;size:100"`
	AddressDetail string    `gorm:"not null;size:255"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserRole struct {
	ID       int    `gorm:"primaryKey"`
	RoleName string `gorm:"not null;size:50;unique;check:role_name IN ('patient','doctor','admin')"`
}

type UserDoctor struct {
	ID             int64  `gorm:"primaryKey"`
	UserID         string `gorm:"not null"`
	User           User   `gorm:"foreignKey:UserID"`
	IsVerified     bool   `gorm:"not null"`
	Specialization string `gorm:"not null;size:100"`
	StrNo          string `gorm:"not null;size:100"`
	BioDesc        string `gorm:"not null;size:255"`
	WorkYears      int    `gorm:"not null"`
	Rating         float64
}

type UserPatient struct {
	ID              int64  `gorm:"primaryKey"`
	UserID          string `gorm:"not null"`
	User            User   `gorm:"foreignKey:UserID"`
	DiabetesHistory bool   `gorm:"not null"`
	DiabetesType    string `gorm:"size:50"`
	DiagnosisDate   time.Time
}
