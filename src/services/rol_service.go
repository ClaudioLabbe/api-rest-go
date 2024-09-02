package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"fmt"
	"log"
)

func firstRol(rol *models.Rol, id int) error {
	result := config.DB.First(&rol, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllRols() ([]models.Rol, error) {
	var rols []models.Rol

	result := config.DB.Find(&rols)

	// Maneja posibles errores
	if result.Error != nil {
		log.Printf("Error getting roles: %v", result.Error)
		return nil, result.Error
	}

	return rols, result.Error
}

func CreateRol(rol models.Rol) (string, error) {
	result := config.DB.Create(&rol)

	// Maneja posibles errores
	if result.Error != nil {
		log.Printf("Error creating role: %v", result.Error)

		return fmt.Sprintf("Error creating role: %v", result.Error), result.Error
	}

	return "Rol created successfully", nil
}

func GetRolById(id int) (*models.Rol, error) {
	var rol models.Rol

	err := firstRol(&rol, id)

	return &rol, err
}

func DeleteRol(id int) error {
	var rols []models.Rol
	var rol models.Rol

	fmt.Println(id)

	err := firstRol(&rol, id)
	if err != nil {
		return err
	}

	result := config.DB.Delete(&rols, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
