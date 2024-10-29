package handlers

import (
	"log"
	"net/http"
	"strconv"

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
		log.Printf("Validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	create := database.DB.Create(&student)
	if create.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create student"})
		return
	}

	log.Printf("Successfully created student: %v", student)
	c.JSON(http.StatusOK, gin.H{"message": "student created successfully"})
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	findall := database.DB.Find(&students)
	if findall.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch students"})
		return
	}
	log.Printf("Successfully get all student: %v", students)
	c.JSON(http.StatusOK, students)
}
func GetStudentById(c *gin.Context) {
	student := models.Student{}
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	
	findId := database.DB.First(&student, id)
	if findId.Error != nil {
		log.Printf("Database error: %v", findId.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "no student with the id found"})	
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch student"})
		return
	}
	log.Printf("Successfully get student: %v", student)
	c.JSON(http.StatusOK, student)
}
func UpdateStudent(c *gin.Context) {
	student := models.Student{}
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	// first retrieve the record
	findId := database.DB.First(&student, id)
	if findId.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no student with the id found"})	
		return
	}

	// bind the update student data
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	// to validate the update student data
	if err := student.Validate(); err != nil {
		log.Printf("Validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := database.DB.Save(&student)
	if update.Error != nil {
		log.Printf("Database error: %v", update.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update student"})
	}

	log.Printf("Successfully update student: %v", student)
	c.JSON(http.StatusOK, student)

}
func DeleteStudent(c *gin.Context) {
	student := models.Student{}
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	findId := database.DB.First(&student, id)
	if findId.Error != nil {
		log.Printf("Validation error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "no student with the id found"})	
		return
	}

	delete := database.DB.Delete(&student)
	if delete.Error != nil {
		log.Printf("Database error: %v", delete.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete student"})
	}
	log.Printf("Successfully delete student: %v", student)
	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}