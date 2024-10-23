package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"fmt"
	"log"
	"time"
)

func PostSaleService(sale models.Sale) (string, error) {
	sale.SaleDate = time.Now()

	for _, sd := range sale.SalesDetails {
		fmt.Println(sd.ProductId)
		var product models.Product

		err := firstProduct(&product, sd.ProductId)
		if err != nil {
			return fmt.Sprintf("Product not found: %v", sd.ProductId), err
		}

		fmt.Println(product)

		if sd.Quantity > product.QuantityInStock {
			errMsg := fmt.Sprintf("Not enough stock for product %d: requested %d, available %d", sd.ProductId, sd.Quantity, product.QuantityInStock)
			return errMsg, fmt.Errorf(errMsg)
		}

		price := product.Price * float64(sd.Quantity)

		if price != sd.Price {
			errMsg := fmt.Sprintf("The price of the product does not match what was sent. Expected total: %f, received total: %f", price, sd.Price)
			return errMsg, fmt.Errorf(errMsg)
		}

		result := config.DB.Create(&sale)

		if result.Error != nil {
			log.Printf("Error creating sale: %v", result.Error)

			return fmt.Sprintf("Error creating sale: %v", result.Error), result.Error
		}

		product.QuantityInStock -= sd.Quantity
		config.DB.Save(&product)
	}

	return "Sale created successfully", nil
}
