package main

import (
	"k9-system/config"
	"k9-system/database"
	"k9-system/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	

	database.Migrate()
	database.SeedAdminUser()
	database.SeedRoles()
	database.SeedSystemModules()
	database.SeedModuleAccess()
	database.SeedAdminUser()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "K9 API Running",
		})
	})

	router.Run(":8080")
}