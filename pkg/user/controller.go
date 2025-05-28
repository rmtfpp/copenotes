package user

import (
	"log"

	"github.com/google/uuid"
	"github.com/rmtfpp/copenotes/pkg/database"
)

func CreateUser(firstname string, lastname string, username string, email string, password string) error {

	id := uuid.New().String()

	user := User{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
		UserName:  username,
		Email:     email,
		Password:  password,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("failed to create user: %v", result.Error)
		return result.Error
	}

	log.Printf("User created successfully: %+v\n", user)

	return nil
}

func DeleteUser(idStr string) error {

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("failed to delate user: %v", err)
		return err
	}

	result := database.DB.Delete(&User{}, id)
	if result.Error != nil {
		log.Printf("failed to delate user: %v", result.Error)
		return result.Error
	}

	log.Printf("User deleted successfully: %+v\n", idStr)
	return nil
}

func UsernameExists(username string) (bool, error) {
	var count int64
	err := database.DB.Model(&User{}).Where("user_name = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func EmailExists(email string) (bool, error) {
	var count int64
	err := database.DB.Model(&User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := database.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		log.Printf("failed to get user by username: %v", err)
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		return nil, err
	}
	return &user, nil
}

func MigrateUsers() {
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate user database: %v", err)
	}
}
