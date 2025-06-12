package main

import (
	"box/internal/core/service"
	"box/internal/core/usecase/sync"
	"box/internal/storage/database"
	"box/pkg/gui"
	"log"
	"os"
)

func main() {
	database.InitDB()

	// 获取当前工作目录作为扫描根路径
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}

	// 扫描文件系统并同步到数据库
	go func() {
		if err := sync.SyncDirectory(rootPath); err != nil {
			log.Printf("文件系统同步失败: %v", err)
		}
	}()

	// 创建服务实例
	boxService := &service.LocalBoxService{}

	// 启动GUI界面并传递服务实例
	gui.StartGUI(boxService)
}
