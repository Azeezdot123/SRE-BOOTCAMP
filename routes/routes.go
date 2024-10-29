package routes

import (
	"github.com/azeezdot123/SRE-BOOTCAMP/handlers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {
	router.POST("/api/v1/students", handlers.CreateStudent)
	// router.GET("/api/v1/students", handlers.GetStudents)
	// router.GET("/api/v1/students/:id", handlers.GetStudent)
	// router.PUT("/api/v1/students/:id", handlers.UpdateStudent)
	// router.DELETE("/api/v1/students/:id", handlers.DeleteStudent)
}

// func HealthCheck(router *gin.Engine) {
// 	router.GET("/api/v1/healthcheck", handlers.HealthCheck)
// }