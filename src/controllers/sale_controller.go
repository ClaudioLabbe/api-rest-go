package controllers

import (
	"api-rest-go/src/models"
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSale(g *gin.Context) {
	var newSale models.Sale

	if err := g.BindJSON(&newSale); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": utils.Invalid_request})
		return
	}

	menssage, err := services.PostSaleService(newSale)

	if err != nil {
		response := utils.Response{
			Message:    menssage,
			Data:       "",
			StatusCode: http.StatusBadRequest,
		}

		g.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.Response{
		Message:    utils.Successful_request,
		Data:       menssage,
		StatusCode: http.StatusOK,
	}
	g.IndentedJSON(http.StatusCreated, response)
}
