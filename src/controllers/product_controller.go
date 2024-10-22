package controllers

import (
	"api-rest-go/src/models"
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProducts(g *gin.Context) {
	products := services.GetAllProducts()

	if products == nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get product"})
		return
	}

	g.IndentedJSON(http.StatusOK, products)
}

func PostProduct(g *gin.Context) {
	var newProduct models.Product

	if err := g.BindJSON(&newProduct); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Invalid_request})
		return
	}

	services.CreateProduct(newProduct)

	g.IndentedJSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func GetProductById(g *gin.Context) {
	idStr := g.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Id_invalid})
		return
	}

	product, err := services.GetProductById(id)
	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Product_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting Product"})
		return
	}

	// Si no hay errores, devuelve el álbum encontrado
	g.JSON(http.StatusOK, product)
}

func UpdateProductById(g *gin.Context) {

	idStr := g.Param("id")

	var updateProduct models.Product

	if err := g.BindJSON(&updateProduct); err != nil {
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

	product, err := services.UpdateProductById(updateProduct, id)

	if err != nil {
		// Si el error es gorm.ErrRecordNotFound, responde con un error 404 Not Found
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Product_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Product"})
		return
	}

	g.IndentedJSON(http.StatusOK, product)
}

func DeleteProduct(g *gin.Context) {
	idStr := g.Param("id")

	// Convertir el string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Responder con un error si la conversión falla
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Id_invalid})
		return
	}

	err = services.DeleteProduct(id)

	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, gin.H{"error": utils.Product_not_found})
			return
		}
		// Para otros errores, responde con un error 500 Internal Server Error
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Product"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
