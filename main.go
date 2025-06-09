package main

import (
	"box/backend/database"
	"box/frontend/gui"
	"box/internal/service"
)

func main() {
	database.InitDB()

	// 创建服务实例
	boxService := &service.LocalBoxService{}
	
	// 启动GUI界面并传递服务实例
	gui.StartGUI(boxService)
}
