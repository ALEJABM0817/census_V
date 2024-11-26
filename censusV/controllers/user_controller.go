package controllers

import (
	"censusV/database"
	"censusV/models"
	"censusV/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Registro de usuario
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hashear contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}
	user.Password = string(hashedPassword)

	// Inicializar preferencias si están vacías
	if user.Preferences == "" {
		user.Preferences = "{}"
	}

	// Guardar usuario
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado correctamente"})
}

// Inicio de sesión
func LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user models.User

	// Obtener credenciales del request
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el usuario existe
	if err := database.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		fmt.Println("Error al consultar el usuario en la db", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		fmt.Println("Error al verificar la contraseña", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar un token JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		fmt.Println("Error al verificar el token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token", "details": err.Error()})
		return
	}

	// Ingreso exitoso
	c.JSON(http.StatusOK, gin.H{
		"message":     "Inicio de sesión exitoso",
		"token":       token,
		"preferences": user.Preferences,
	})
}

func SavePreferences(c *gin.Context) {
	var preferences map[string]interface{}
	if err := c.ShouldBindJSON(&preferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Obtener el usuario autenticado
	userID := c.MustGet("userID").(uint)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Convertir preferencias a JSON
	preferencesJSON, err := json.Marshal(preferences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron procesar las preferencias"})
		return
	}

	// Guardar preferencias
	user.Preferences = string(preferencesJSON)
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron guardar las preferencias"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preferencias guardadas correctamente"})
}

func CheckAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"isAuthenticated": false, "message": "No autorizado. Por favor, inicia sesión."})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"isAuthenticated": false, "message": "Token inválido.", "error": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"isAuthenticated": false, "message": "Token inválido: user_id no encontrado."})
			return
		}
		var user models.User
		if err := database.DB.First(&user, uint(userID)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"isAuthenticated": false, "message": "Usuario no encontrado."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"isAuthenticated": true, "user": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"isAuthenticated": false, "message": "Token inválido."})
	}
}
