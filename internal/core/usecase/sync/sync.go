package sync

import (
	"box/internal/storage/repos"
	"box/internal/storage/models"
	"os"
	"path/filepath"
)

// SyncDirectory 同步指定目录下的所有文件和目录到数据库
func SyncDirectory(rootPath string) error {
	// 确保根目录Box存在
	if _, err := ensureBox(rootPath, nil); err != nil {
		return err
	}

	// 递归遍历目录
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过根目录自身
		if path == rootPath {
			return nil
		}

		// 处理目录
		if info.IsDir() {
			parentPath := filepath.Dir(path)
			parentBox, err := getOrCreateBox(parentPath)
			if err != nil {
				return err
			}
			_, err = ensureBox(path, parentBox)
			return err
		}

		// 处理文件
		parentPath := filepath.Dir(path)
		parentBox, err := getOrCreateBox(parentPath)
		if err != nil {
			return err
		}
		return createFile(path, parentBox)
	})
}

// ensureBox 确保目录对应的Box存在
func ensureBox(dirPath string, parentBox *models.Box) (*models.Box, error) {
	box, err := crud.GetBoxByName(filepath.Base(dirPath))
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

	if err := crud.CreateBox(newBox); err != nil {
		return nil, err
	}
	return newBox, nil
}

// getOrCreateBox 获取或创建目录对应的Box
func getOrCreateBox(dirPath string) (*models.Box, error) {
	box, err := crud.GetBoxByName(filepath.Base(dirPath))
	if err == nil {
		return box, nil
	}
	return ensureBox(dirPath, nil)
}

// createFile 创建文件记录并关联到Box
func createFile(filePath string, parentBox *models.Box) error {
	// 检查文件是否已存在
	_, err := crud.GetFileByPath(filePath)
	if err == nil {
		return nil // 文件已存在
	}

	// 创建新文件记录
	file := &models.File{
		Name: filepath.Base(filePath),
		Path: filePath,
	}
	if err := crud.CreateFile(file); err != nil {
		return err
	}

	// 关联到父Box
	return crud.AddBoxFile(parentBox.ID, int(file.ID))
}
