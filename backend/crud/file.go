package crud

import (
	"box/backend/database"
	"box/backend/models"
)

// CreateFile 创建文件记录
func CreateFile(file *models.File) error {
	return database.DB.Create(file).Error
}

// GetFileByID 根据ID获取文件
func GetFileByID(id int) (*models.File, error) {
	var file models.File
	err := database.DB.Preload("Tags").Preload("Boxes").First(&file, id).Error
	return &file, err
}

// UpdateFile 更新文件信息
func UpdateFile(file *models.File) error {
	return database.DB.Save(file).Error
}

// DeleteFile 删除文件
func DeleteFile(id int) error {
	return database.DB.Delete(&models.File{}, id).Error
}

// SetFileTags 设置文件的标签关联
func SetFileTags(fileID int, tagIDs []int) error {
	file := models.File{ID: fileID}
	return database.DB.Model(&file).Association("Tags").Replace(tagIDs)
}

// SetFileBoxes 设置文件的Box关联
func SetFileBoxes(fileID int, boxIDs []int) error {
	file := models.File{ID: fileID}
	return database.DB.Model(&file).Association("Boxes").Replace(boxIDs)
}

// GetFileByPath 根据路径获取文件
func GetFileByPath(path string) (*models.File, error) {
	var file models.File
	err := database.DB.Where("path = ?", path).First(&file).Error
	return &file, err
}
