package routes

import (
	"github.com/azeezdot123/SRE-BOOTCAMP/handlers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	r.POST("/api/v1/students", handlers.CreateStudent)
	r.GET("/api/v1/students", handlers.GetStudents)
	r.GET("/api/v1/student/:id", handlers.GetStudentById)
	r.PUT("/api/v1/student/:id", handlers.UpdateStudent)
	r.DELETE("/api/v1/student/:id", handlers.DeleteStudent)
}

func HealthCheck(r *gin.Engine) {
	r.GET("/api/v1/healthcheck", handlers.HealthCheck)
}