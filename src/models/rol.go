package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model

	//ID     uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name   string `json:"name" gorm:"not null;unique"`
	Status bool   `json:"status" gorm:"not null"`
}

func (Rol) TableName() string {
	return "rol"
}
