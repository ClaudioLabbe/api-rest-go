package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-rest-go/src/models"
)

func GetAllAlbum(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, models.Albums)
}

func PostAlbum(g *gin.Context) {
	var newAlbum models.Album

	if err := g.BindJSON(&newAlbum); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	models.Albums = append(models.Albums, newAlbum)

	g.IndentedJSON(http.StatusCreated, models.Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	var album models.Album

	for _, a := range models.Albums {
		if a.ID == id {
			album = a
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such album"})
}

func UpdateAlbumById(c *gin.Context) {
	id := c.Param("id")

	var album models.Album

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	for i := range models.Albums {
		if models.Albums[i].ID == id {
			models.Albums[i].Artist = album.Artist
			models.Albums[i].Title = album.Title
			models.Albums[i].Year = album.Year

			c.IndentedJSON(http.StatusOK, models.Albums)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})

}

func DeleteAlbum(g *gin.Context) {
	id := g.Param("id")

	fmt.Println(id)

	for i := range models.Albums {
		if models.Albums[i].ID == id {
			models.Albums = append(models.Albums[:i], models.Albums[i+1:]...)
			g.IndentedJSON(http.StatusCreated, models.Albums)
			return
		}
	}

	g.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}
