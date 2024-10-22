package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"log"
)

func GetAllProducts() []models.Product {
	var products []models.Product

	result := config.DB.Find(&products)

	if result.Error != nil {
		log.Printf("Error al obtener los productos: %v", result.Error)
		return nil
	}

	return products
}

func firstProduct(product *models.Product, id int) error {
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetProductById(id int) (*models.Product, error) {
	var product models.Product

	err := firstProduct(&product, id)

	return &product, err
}

func CreateProduct(newProduct models.Product) {
	config.DB.Create(&newProduct)
}

func UpdateProductById(updateProduct models.Product, id int) (*models.Product, error) {

	var product models.Product

	err := firstProduct(&product, id)
	if err != nil {
		return nil, err
	}

	result := config.DB.Model(&product).Updates(updateProduct)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func DeleteProduct(id int) error {
	var products []models.Product
	var product models.Product

	err := firstProduct(&product, id)
	if err != nil {
		return err
	}

	result := config.DB.Delete(&products, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
