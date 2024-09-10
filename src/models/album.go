package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null;unique"`
	Artist string `json:"artist" gorm:"not null"`
	Year   int    `json:"year" gorm:"not null"`
}

func (Album) TableName() string {
	return "album"
}
