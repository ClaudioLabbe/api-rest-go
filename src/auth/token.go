package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Clave secreta utilizada para firmar el token
var jwtKey = []byte(os.Getenv("API_SECRET"))

// Estructura de las Claims que queremos incluir en el JWT
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerarToken genera un JWT para un usuario dado
func GenerarToken(email string) (string, error) {
	// Definir las claims
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expira en 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Crear el token con el método de firma HS256 y las claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerificarToken verifica un token JWT y extrae las claims
func VerificarToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	// Parsear el token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Validar el token y las claims
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
