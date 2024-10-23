package models

import "gorm.io/gorm"

type SalesDetails struct {
	gorm.Model
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	SaleId    int     `json:"sale_id" gorm:"not null"`
	ProductId int     `json:"product_id" gorm:"not null"`
}

func (SalesDetails) TableName() string {
	return "salesdetails"
}
