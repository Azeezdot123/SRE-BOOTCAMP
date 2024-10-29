package handlers

import (
	"log"
	"net/http"

	"github.com/azeezdot123/SRE-BOOTCAMP/database"
	"github.com/azeezdot123/SRE-BOOTCAMP/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	student := models.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	// This validate to check if the required fields are present
	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	create := database.DB.Create(&student)
	if create.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create student"})
		return
	}
	c.JSON(http.StatusCreated, student)
	c.JSON(http.StatusOK, student)
}

// func GetStudents(c *gin.Context) {}
// func GetStudent(c *gin.Context) {}
// func UpdateStudent(c *gin.Context) {}
// func DeleteStudent(c *gin.Context) {}