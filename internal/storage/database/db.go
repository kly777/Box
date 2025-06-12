package database

import (
	"box/internal/storage/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

var DB *gorm.DB

func InitDB() {
	once.Do(initDB)
}

func initDB() {
	db, err := gorm.Open(sqlite.Open(".db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Box{}, &models.File{}, &models.Tag{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	DB = db // 将数据库连接赋值给全局变量
}
