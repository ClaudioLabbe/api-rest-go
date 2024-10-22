package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string  `json:"name" gorm:"not null;unique"`
	Description     string  `json:"description" gorm:"not null"`
	Price           float64 `json:"price" gorm:"not null"`
	QuantityInStock int     `json:"quantity_in_stock" gorm:"not null"`
	CategoryId      uint64  `json:"category_id" gorm:"not null"`
}

func (Product) TableName() string {
	return "products"
}
