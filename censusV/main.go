package main

import (
	"log"
	"os"
	"time"

	"censusV/database"
	"censusV/loaders"
	"censusV/routes"

	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var loadDataOnRun bool = false

func main() {

	// Inicializar la base de datos
	log.Println("Inicializando la base de datos...")
	database.InitDB()

	if loadDataOnRun {
		// Proporciona la ruta completa al archivo source_data.csv
		const filePath = "./data/source_data.csv"
		// read a file
		data := loaders.ReadData(filePath)
		// Se intenta guardar en la base de datos
		if err := loaders.SaveDataIntoDB(data); err != nil {
			log.Println("error")
			return
		}
	}
	// Fin de carga de datos ------------------------------------------------------------------------------------------------

	// Configurar el router
	router := routes.SetupRouter()

	// Obtener el puerto del entorno o usar el puerto 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Iniciar el servidor
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
