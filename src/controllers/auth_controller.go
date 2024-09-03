package controllers

import (
	"api-rest-go/src/services"
	"api-rest-go/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controlador para el login que genera un token JWT
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println(credentials)

	user, err := services.GetUserByEmail(credentials.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	response := utils.VerifyPassword(user.Password, credentials.Password)

	// Validación ficticia de usuario y contraseña
	if response != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Genera el token JWT
	token, err := utils.GenerateJWT(credentials.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
