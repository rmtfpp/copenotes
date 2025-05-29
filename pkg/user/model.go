package user

import (
	"time"
)

type User struct {
	ID        string   `gorm:"unique;type:text;primaryKey"`
	FirstName string   `gorm:"not null"`
	LastName  string   `gorm:"not null"`
	UserName  string   `gorm:"not null"`
	Email     string   `gorm:"unique;not null"`
	Password  string   `gorm:"not null"`
	Session   *Session `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Session struct {
	UserID       string    `gorm:"type:text;primaryKey"`
	SessionToken string    `gorm:"not null"`
	CSRFToken    string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	User         User      `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
