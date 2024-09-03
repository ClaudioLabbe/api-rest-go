package controllers

import (
	"api-rest-go/src/models"
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRols(g *gin.Context) {

	rols, err := services.GetAllRols()

	if rols == nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusInternalServerError,
		}

		g.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.Response{
		Message:    utils.Successful_request,
		Data:       rols,
		StatusCode: http.StatusOK,
	}

	g.IndentedJSON(http.StatusOK, response)
}

func PostRol(g *gin.Context) {
	var rol models.Rol

	if err := g.BindJSON(&rol); err != nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusPaymentRequired,
		}

		g.IndentedJSON(http.StatusPaymentRequired, response)
		return
	}
	if rol.Name == "" {
		response := utils.Response{
			Message:    utils.Payload_is_required,
			Data:       "",
			StatusCode: http.StatusPaymentRequired,
		}

		g.IndentedJSON(http.StatusPaymentRequired, response)
		return
	}

	menssage, err := services.CreateRol(rol)

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

func GetRolById(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusInternalServerError,
		}

		g.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	rol, err := services.GetRolById(id)

	if err != nil {
		fmt.Println(err.Error())
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusNotFound,
		}

		g.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := utils.Response{
		Message:    utils.Successful_request,
		Data:       rol,
		StatusCode: http.StatusOK,
	}
	g.IndentedJSON(http.StatusOK, response)
}

func DeleteRol(g *gin.Context) {
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusInternalServerError,
		}

		g.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	err = services.DeleteRol(id)

	if err != nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusNotFound,
		}

		g.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := utils.Response{
		Message:    utils.Successful_request,
		Data:       "",
		StatusCode: http.StatusOK,
	}

	g.IndentedJSON(http.StatusOK, response)
}
