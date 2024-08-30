package controllers

import (
	"api-rest-go/src/services"
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
		g.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los álbumes"})
		return
	}

	g.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(g *gin.Context) {
	var newAlbum models.Album

	if err := g.BindJSON(&newAlbum); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	services.CreateAlbum(newAlbum)

	g.IndentedJSON(http.StatusCreated, gin.H{"message": "Album created successfully"})
}

func GetAlbumById(c *gin.Context) {
	idStr := c.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	album, err := services.GetAlbumById(id)
	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Álbum no encontrado"})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el álbum"})
		return
	}

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such album"})
	// Si no hay errores, devuelve el álbum encontrado
	c.JSON(http.StatusOK, album)
}

func UpdateAlbumById(c *gin.Context) {

	idStr := c.Param("id")

	var updateAlbum models.Album

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	album, err := services.UpdateAlbumById(updateAlbum, id)

	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Álbum no encontrado"})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el álbum"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(g *gin.Context) {
	idStr := g.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = services.DeleteAlbum(id)

	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": "Álbum no encontrado"})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el álbum"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Album eliminado correctamente"})
}
