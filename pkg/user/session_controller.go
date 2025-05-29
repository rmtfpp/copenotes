package user

import (
	"errors"
	"log"
	"time"

	"github.com/rmtfpp/copenotes/pkg/database"
	"gorm.io/gorm"
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

func GetSession(id string) (*Session, error) {
	var session Session
	err := database.DB.Where("user_id = ?", id).First(&session).Error
	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		return nil, err
	}
	return &session, nil
}

func GetSessionToken(id string) (string, error) {
	session, _ := GetSession(id)
	return session.SessionToken, nil
}

func GetCSRFToken(id string) (string, error) {
	session, _ := GetSession(id)
	return session.CSRFToken, nil
}

func DeleteSession(id string) error {
	var session Session
	err := database.DB.Where("user_id = ?", id).Delete(&session).Error
	if err != nil {
		log.Printf("failed to delete user by id: %v", err)
		return err
	}
	return nil
}

func HasSession(id string) (bool, error) {
	var session Session
	err := database.DB.Where("user_id = ?", id).First(&session).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		log.Printf("failed to delete user by id: %v", err)
		return false, err
	}
	return true, nil
}

func MigrateSessions() {
	if err := database.DB.AutoMigrate(&Session{}); err != nil {
		log.Fatalf("failed to migrate session database: %v", err)
	}
}
