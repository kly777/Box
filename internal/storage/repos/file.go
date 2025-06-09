package repos

import (
	"box/internal/storage/database"
	"box/internal/storage/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

// CreateFile 创建文件记录
func CreateFile(file *models.File) error {
	log.Printf("[FileRepo] 创建文件记录: %s", file.Path)
	err := database.DB.Create(file).Error
	if err != nil {
		log.Printf("[FileRepo] 创建文件失败: %s, 错误: %v", file.Path, err)
	}
	return err
}

// GetFileByID 根据ID获取文件
func GetFileByID(id int) (*models.File, error) {
	var file models.File
	err := database.DB.Preload("Tags").Preload("Boxes").First(&file, id).Error
	return &file, err
}

// UpdateFile 更新文件信息
func UpdateFile(file *models.File) error {
	log.Printf("[FileRepo] 更新文件记录(ID:%d): %s", file.ID, file.Path)
	err := database.DB.Save(file).Error
	if err != nil {
		log.Printf("[FileRepo] 更新文件失败(ID:%d): %v", file.ID, err)
	}
	return err
}

// DeleteFile 删除文件
func DeleteFile(id int) error {
	log.Printf("[FileRepo] 删除文件记录(ID:%d)", id)
	err := database.DB.Delete(&models.File{}, id).Error
	if err != nil {
		log.Printf("[FileRepo] 删除文件失败(ID:%d): %v", id, err)
	}
	return err
}

// SetFileTags 设置文件的标签关联
func SetFileTags(fileID int, tagIDs []int) error {
	log.Printf("[FileRepo] 设置文件标签关联(ID:%d Tags:%v)", fileID, tagIDs)
	file := models.File{ID: fileID}
	err := database.DB.Model(&file).Association("Tags").Replace(tagIDs)
	if err != nil {
		log.Printf("[FileRepo] 设置标签失败(ID:%d): %v", fileID, err)
	}
	return err
}

// SetFileBoxes 设置文件的Box关联
func SetFileBoxes(fileID int, boxIDs []int) error {
	log.Printf("[FileRepo] 设置文件Box关联(ID:%d Boxes:%v)", fileID, boxIDs)
	file := models.File{ID: fileID}
	err := database.DB.Model(&file).Association("Boxes").Replace(boxIDs)
	if err != nil {
		log.Printf("[FileRepo] 设置Box失败(ID:%d): %v", fileID, err)
	}
	return err
}

// GetFileByPath 根据路径获取文件
func GetFileByPath(path string) (*models.File, error) {
	var file models.File
	log.Printf("[FileRepo] 按路径查询文件: %s", path)
	err := database.DB.Where("path = ?", path).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("[FileRepo] 文件不存在: %s", path)
		return nil, nil
	}
	return &file, err
}

// 根据Box ID获取文件列表
func GetFilesByBoxID(boxID uint) ([]models.File, error) {
	var files []models.File
	err := database.DB.Joins("JOIN file_boxes ON file_boxes.file_id = files.id").
		Where("file_boxes.box_id = ?", boxID).
		Find(&files).Error
	return files, err
}
