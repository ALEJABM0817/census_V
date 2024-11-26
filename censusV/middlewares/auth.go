package middlewares

import (
	"censusV/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta token de autorización"})
			c.Abort()
			return
		}

		// Validar el token
		token, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Extraer el userID del token
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["sub"].(float64))
		c.Set("userID", userID)

		c.Next()
	}
}
