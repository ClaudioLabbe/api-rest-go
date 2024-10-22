package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`
}

func (Category) TableName() string {
	return "categories"
}
