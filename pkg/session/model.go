package session

import "time"

type Session struct {
	UserID       string    `gorm:"unique;type:text;primaryKey"`
	SessionToken string    `gorm:"not null"`
	CSRFToken    string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
