package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"fmt"
	"log"
)

func firstUser(user *models.User, id int) error {
	result := config.DB.Preload("Rol").First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := config.DB.Preload("Rol").Find(&users)

	// Maneja posibles errores
	if result.Error != nil {
		log.Printf("Error getting users: %v", result.Error)
		return nil, result.Error
	}

	return users, result.Error
}

func CreateUser(user models.User) (string, error) {
	result := config.DB.Create(&user)

	// Maneja posibles errores
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)

		return fmt.Sprintf("Error creating user: %v", result.Error), result.Error
	}

	return "User created successfully", nil
}

func GetUserById(id int) (*models.User, error) {
	var user models.User

	err := firstUser(&user, id)

	return &user, err
}

func DeleteUser(id int) error {
	var users []models.User
	var user models.User

	err := firstUser(&user, id)
	if err != nil {
		return err
	}

	result := config.DB.Delete(&users, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := config.DB.Preload("Rol").Where("email = ?", email).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, err.Error
}
