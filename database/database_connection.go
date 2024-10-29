package database

import (
	"github.com/azeezdot123/SRE-BOOTCAMP/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database connection
func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Student{})
	return db
}

var Client *gorm.DB = DatabaseConnection()
