package middleware

import (
	"api-rest-go/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtiene el token del header Authorization
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// El formato debe ser "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// Valida el token
		claims, err := utils.ValidateJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !isAuthorized(claims, c.FullPath()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
			c.Abort()
			return
		}

		// Establece las claims en el contexto
		c.Set("email", claims.Email)
		c.Next()
	}
}

func isAuthorized(claims *utils.Claims, path string) bool {

	switch claims.Rol {
	case "User":
		if strings.HasPrefix(path, "/album/") {
			return true
		}
	case "Owner":
		return true
	case "Admin":
		return true
	default:
		return false
	}

	return false
}
