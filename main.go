package main

import (
	"log"

	"github.com/azeezdot123/SRE-BOOTCAMP/database"
	"github.com/azeezdot123/SRE-BOOTCAMP/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	
	// Initialize the DB
	database.InitDB()

	// This is where we define the API routes
	routes.StudentRoutes(router)

	// routes.HealthCheck(router)
	routes.HealthCheck(router)

	router.Run()
}
