package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

// Claims define la estructura de las claims del JWT
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Genera un nuevo token JWT para un usuario
func GenerateJWT(email string) (string, error) {
	// Define las claims
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Expira en 15 minutos
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
