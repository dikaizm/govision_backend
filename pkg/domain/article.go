package domain

import "time"

type Article struct {
	ID        string    `gorm:"primaryKey"`
	Title     string    `gorm:"not null;size:255"`
	Body      string    `gorm:"not null"`
	Image     string    `gorm:"not null;type:text"`
	AuthorID  string    `gorm:"not null;size:255"`
	Author    User      `gorm:"foreignKey:AuthorID;references:ID"`
	ReadCount int       `gorm:"not null;default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
