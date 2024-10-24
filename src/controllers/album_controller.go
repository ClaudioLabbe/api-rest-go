package controllers

import (
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"fmt"
	"net/http"
	"strconv"

	"api-rest-go/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllAlbum(g *gin.Context) {

	albums := services.GetAllAlbum()

	// Verifica si hay datos para devolver
	if albums == nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get albums"})
		return
	}

	g.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(g *gin.Context) {
	var newAlbum models.Album

	if err := g.BindJSON(&newAlbum); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Invalid_request})
		return
	}

	services.CreateAlbum(newAlbum)

	g.IndentedJSON(http.StatusCreated, gin.H{"message": "Album created successfully"})
}

func GetAlbumById(g *gin.Context) {
	idStr := g.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Id_invalid})
		return
	}

	album, err := services.GetAlbumById(id)
	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Album_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting album"})
		return
	}

	// Si no hay errores, devuelve el álbum encontrado
	g.JSON(http.StatusOK, album)
}

func UpdateAlbumById(g *gin.Context) {

	idStr := g.Param("id")

	var updateAlbum models.Album

	if err := g.BindJSON(&updateAlbum); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Invalid_request})
		return
	}

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Id_invalid})
		return
	}

	album, err := services.UpdateAlbumById(updateAlbum, id)

	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Album_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating album"})
		return
	}

	g.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(g *gin.Context) {
	idStr := g.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Id_invalid})
		return
	}

	err = services.DeleteAlbum(id)

	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Album_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting album"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
