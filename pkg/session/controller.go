package session

import (
	"log"
	"time"

	"github.com/rmtfpp/copenotes/pkg/database"
)

func CreateSession(id string, sessionToken string, csrfToken string) error {

	session := Session{
		UserID:       id,
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}

	result := database.DB.Create(&session)
	if result.Error != nil {
		log.Printf("failed to create user: %v", result.Error)
		return result.Error
	}
	return nil
}

func MigrateSessions() {
	if err := database.DB.AutoMigrate(&Session{}); err != nil {
		log.Fatalf("failed to migrate session database: %v", err)
	}
}

func GetSessionToken(uuid string) (string, error) {
	var session Session
	err := database.DB.Where("user_id = ?", uuid).First(&session).Error
	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		return "", err
	}
	return session.SessionToken, nil
}

func GetCSRFToken(uuid string) (string, error) {
	var session Session
	err := database.DB.Where("user_id = ?", uuid).First(&session).Error
	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		return "", err
	}
	return session.CSRFToken, nil
}
