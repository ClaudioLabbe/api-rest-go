package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"log"
)

func GetAllCategories() []models.Category {
	var categories []models.Category

	result := config.DB.Find(&categories)

	if result.Error != nil {
		log.Printf("Error al obtener las categorias: %v", result.Error)
		return nil
	}

	return categories
}
