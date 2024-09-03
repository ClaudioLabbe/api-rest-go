package models

import (
	"api-rest-go/src/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	//Id       uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Lastname string `json:"lastname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	RolId    uint64 `json:"rol_id"`
	Rol      Rol    `json:"rol"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	passwordHashed, err := utils.Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(passwordHashed)

	return nil
}
