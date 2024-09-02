package controllers

import (
	"api-rest-go/src/models"
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(g *gin.Context) {

	users, err := services.GetAllUsers()

	if users == nil {
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
		Data:       users,
		StatusCode: http.StatusOK,
	}

	g.IndentedJSON(http.StatusOK, response)
}

func PostUser(g *gin.Context) {
	var user models.User

	if err := g.BindJSON(&user); err != nil {
		response := utils.Response{
			Message:    err.Error(),
			Data:       "",
			StatusCode: http.StatusPaymentRequired,
		}

		g.IndentedJSON(http.StatusPaymentRequired, response)
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Lastname == "" || user.RolId == 0 {
		response := utils.Response{
			Message:    utils.Payload_is_required,
			Data:       "",
			StatusCode: http.StatusPaymentRequired,
		}

		g.IndentedJSON(http.StatusPaymentRequired, response)
		return
	}

	menssage, err := services.CreateUser(user)

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

func GetUserById(g *gin.Context) {
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

	user, err := services.GetUserById(id)

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
		Data:       user,
		StatusCode: http.StatusOK,
	}
	g.IndentedJSON(http.StatusOK, response)
}

func DeleteUser(g *gin.Context) {
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

	err = services.DeleteUser(id)

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
