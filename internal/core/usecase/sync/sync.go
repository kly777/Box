package sync

import (
	"box/internal/storage/models"
	repos "box/internal/storage/repos"
	"log"
	"os"
	"path/filepath"
)

// SyncDirectory 同步指定目录下的所有文件和目录到数据库
func SyncDirectory(rootPath string) error {
	log.Printf("开始同步目录: %s", rootPath)
	// 确保根目录Box存在
	if _, err := ensureBox(rootPath, nil); err != nil {
		log.Printf("初始化根目录Box失败: %v", err)
		return err
	}

	// 使用队列实现广度优先遍历
	queue := []string{rootPath}
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]

		// 读取当前目录内容
		entries, err := os.ReadDir(currentPath)
		if err != nil {
			log.Printf("读取目录失败: %s, 错误: %v", currentPath, err)
			continue
		}

		// 先处理目录
		for _, entry := range entries {
			path := filepath.Join(currentPath, entry.Name())

			if entry.IsDir() {
				// 处理目录
				parentPath := filepath.Dir(path)
				parentBox, err := getOrCreateBox(parentPath)
				if err != nil {
					log.Printf("获取父Box失败: %s, 错误: %v", parentPath, err)
					continue
				}
				if _, err = ensureBox(path, parentBox); err != nil {
					log.Printf("创建Box失败: %s", path)
					continue
				}
				queue = append(queue, path) // 将子目录加入队列
			}
		}

		// 再处理文件
		for _, entry := range entries {
			path := filepath.Join(currentPath, entry.Name())

			if !entry.IsDir() {
				// 处理文件
				parentPath := filepath.Dir(path)
				parentBox, err := getOrCreateBox(parentPath)
				if err != nil {
					log.Printf("获取文件父Box失败: %s, 错误: %v", parentPath, err)
					continue
				}
				if err := createFile(path, parentBox); err == nil {
					log.Printf("成功创建文件记录: %s", path)
				}
			}
		}
	}

	return nil
}

// ensureBox 确保目录对应的Box存在
func ensureBox(dirPath string, parentBox *models.Box) (*models.Box, error) {
	box, err := repos.GetBoxByName(filepath.Base(dirPath))
	if err == nil {
		return box, nil
	}

	// 创建新Box
	newBox := &models.Box{
		Name: filepath.Base(dirPath),
	}
	if parentBox != nil {
		newBox.ParentID = &parentBox.ID
	}

	if err := repos.CreateBox(newBox); err != nil {
		return nil, err
	}
	return newBox, nil
}

// getOrCreateBox 获取或创建目录对应的Box
func getOrCreateBox(dirPath string) (*models.Box, error) {
	box, err := repos.GetBoxByName(filepath.Base(dirPath))
	if err == nil {
		return box, nil
	}
	return ensureBox(dirPath, nil)
}

// createFile 创建文件记录并关联到Box
func createFile(filePath string, parentBox *models.Box) error {
	// 检查文件是否已存在
	_, err := repos.GetFileByPath(filePath)
	if err == nil {
		return nil // 文件已存在
	}

	// 创建新文件记录
	file := &models.File{
		Name: filepath.Base(filePath),
		Path: filePath,
	}
	if err := repos.CreateFile(file); err != nil {
		return err
	}

	// 关联到父Box
	return repos.AddBoxFile(parentBox.ID, int(file.ID))
}
