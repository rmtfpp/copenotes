package user

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/rmtfpp/copenotes/pkg/database"
)

func CreateFile(userid string, filename string) error {

	user, _ := GetUserById(userid)
	id := uuid.New().String()
	path := strings.ToLower(user.UserName + "/" + filename + ".pdf")

	file := File{
		ID:     id,
		Name:   filename,
		Path:   path,
		UserID: user.ID,
	}

	result := database.DB.Create(&file)
	if result.Error != nil {
		log.Printf("failed to create file: %v", result.Error)
		return result.Error
	}

	log.Printf("file created succesfully\n")
	return nil
}

func MigrateFiles() {
	if err := database.DB.AutoMigrate(&File{}); err != nil {
		log.Fatalf("failed to migrate file database: %v", err)
	}
}
