package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if len(jwtKey) == 0 {
		return "", fmt.Errorf("JWT_SECRET_KEY no está configurado")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("error al firmar el token: %v", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if len(jwtKey) == 0 {
		return nil, fmt.Errorf("JWT_SECRET_KEY no está configurado")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
}
