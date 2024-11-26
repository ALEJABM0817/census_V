package routes

import (
	"censusV/controllers"
	"censusV/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/api/auth/check", controllers.CheckAuth)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Census Data API!",
		})
	})

	userRoutes := router.Group("/api")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}

	censusRoutes := router.Group("/api")
	{
		censusRoutes.GET("/census", controllers.GetCensusData)
		censusRoutes.GET("/summary", controllers.GetSummary)
		censusRoutes.GET("/export", controllers.ExportData)
		//censusRoutes.GET("/visuals", controllers.GetVisuals)
		authRoutes := router.Group("/api/user")
		authRoutes.Use(middlewares.AuthMiddleware())
		{
			authRoutes.POST("/preferences", controllers.SavePreferences)
		}
	}

	return router
}
