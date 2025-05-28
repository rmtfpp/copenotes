package user

import (
	"time"
)

type User struct {
	ID        string `gorm:"unique;type:text;primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	UserName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
