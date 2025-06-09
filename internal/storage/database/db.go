package database

import (
	"box/internal/storage/models"
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
	DB = db // 将数据库连接赋值给全局变量
	fmt.Println()
}
