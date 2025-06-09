package database

import (
	"box/internal/storage/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open(sqlite.Open(".db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&models.Box{}, &models.File{}, &models.Tag{}); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	DB = db // 将数据库连接赋值给全局变量

}
