package handlers

import (
	"net/http"

	"github.com/azeezdot123/SRE-BOOTCAMP/database"
	"github.com/azeezdot123/SRE-BOOTCAMP/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	student := models.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Client.Create(&student)
	c.JSON(http.StatusCreated, student)
}

// func GetStudents(c *gin.Context) {}
// func GetStudent(c *gin.Context) {}
// func UpdateStudent(c *gin.Context) {}
// func DeleteStudent(c *gin.Context) {}