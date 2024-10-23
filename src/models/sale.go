package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Name         string         `json:"name" gorm:"not null"`
	Total        float64        `json:"total" gorm:"not null"`
	SaleDate     time.Time      `json:"sale_date"`
	SalesDetails []SalesDetails `json:"salesDetails" gorm:"foreignKey:SaleId"`
}

func (Sale) TableName() string {
	return "sales"
}
