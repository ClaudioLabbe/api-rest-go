package services

import (
	"api-rest-go/src/config"
	"api-rest-go/src/models"
	"log"
)

func firstAlbum(album *models.Album, id int) error {
	result := config.DB.First(&album, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAlbumById(id int) (*models.Album, error) {
	var album models.Album

	err := firstAlbum(&album, id)

	return &album, err
}

func GetAllAlbum() []models.Album {
	var albums []models.Album

	// Realiza la consulta a la base de datos
	result := config.DB.Find(&albums)

	// Maneja posibles errores
	if result.Error != nil {
		log.Printf("Error al obtener los álbumes: %v", result.Error)
		return nil
	}

	return albums
}

func CreateAlbum(newAlbum models.Album) {
	config.DB.Create(&newAlbum)

}

func UpdateAlbumById(updateAlbum models.Album, id int) (*models.Album, error) {

	var album models.Album

	err := firstAlbum(&album, id)
	if err != nil {
		return nil, err
	}

	result := config.DB.Model(&album).Updates(updateAlbum)

	if result.Error != nil {
		return nil, result.Error
	}

	return &album, nil
}

func DeleteAlbum(id int) error {
	var albums []models.Album
	var album models.Album

	err := firstAlbum(&album, id)
	if err != nil {
		return err
	}

	result := config.DB.Delete(&albums, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
