package database

import (
	"github.com/azeezdot123/SRE-BOOTCAMP/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Database connection
func InitDB() {
        var err error
        DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }

        DB.AutoMigrate(&models.Student{})
	}