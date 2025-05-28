package session

import (
	"log"
	"time"

	"github.com/rmtfpp/copenotes/pkg/database"
	"github.com/rmtfpp/copenotes/pkg/user"
)

func CreateSession(u user.User, sessionToken string, csrfToken string) error {

	session := Session{
		UserID:       u.ID,
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
