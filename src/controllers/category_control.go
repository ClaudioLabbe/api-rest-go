package controllers

import (
	"api-rest-go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllcategories(g *gin.Context) {
	categories := services.GetAllCategories()

	if categories == nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get categories"})
		return
	}

	g.IndentedJSON(http.StatusOK, categories)
}
