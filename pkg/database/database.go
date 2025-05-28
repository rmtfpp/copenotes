package database

import (
	"log"

	"github.com/rmtfpp/copenotes/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {

	connection := config.GetConfig().Database.Connection
	var err error

	DB, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

}
