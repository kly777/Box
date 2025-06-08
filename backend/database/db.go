package database

import (
	"box/backend/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	fmt.Println("InitDB")
	db, err := gorm.Open(sqlite.Open(".db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Box{}, &models.File{}, &models.Tag{})
	fmt.Println()
}
