package main

import (
	"log"

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
	

	router := gin.New()
	router.Use(gin.Logger())

	// This is where we are calling the routes
	routes.StudentRoutes(router)
	// routes.HealthCheck(router)

	

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
