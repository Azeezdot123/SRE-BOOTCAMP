package handlers

import (
	"net/http"
	"runtime"

	"github.com/azeezdot123/SRE-BOOTCAMP/database"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
    // Check the database connection
    db, err := database.DB.DB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy", "error": "Unable to connect to the database"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "message": "The DB is connected and it pings with the application"})

    if err := db.Ping(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy", "error": "Unable to ping the database"})
        return
    }
	var m runtime.MemStats
    runtime.ReadMemStats(&m)

        c.JSON(http.StatusOK, gin.H{
                "status": "OK",
                "message": "The API is healthy.",
                "memory_usage": m.Alloc / 1024 / 1024, // In MB
                "num_goroutines": runtime.NumGoroutine(),
        })
    
}