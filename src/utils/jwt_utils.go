package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("API_SECRET"))

// Claims define la estructura de las claims del JWT
type Claims struct {
	Email string `json:"email"`
	Rol   string `json:"rol"`
	jwt.RegisteredClaims
}

// Genera un nuevo token JWT para un usuario
func GenerateJWT(email string, rol string) (string, error) {

	expirationMinutesStr := os.Getenv("EXPIRATION_MINUTES")
	expirationMinutes, err := strconv.Atoi(expirationMinutesStr)
	if err != nil {
		return "", fmt.Errorf("invalid JWT_EXPIRATION_MINUTES value: %v", err)
	}

	// Define las claims
	claims := &Claims{
		Email: email,
		Rol:   rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMinutes) * time.Minute)), // Expira en 15 minutos
		},
	}

	// Crea el token con el método de firma y claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token con la clave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verifica y parsea el token JWT
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parsea el token y valida con la clave secreta
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
