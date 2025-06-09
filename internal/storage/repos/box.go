package crud

import (
	"box/internal/storage/database"
	"box/internal/storage/models"
)

// CreateBox 创建Box记录
func CreateBox(box *models.Box) error {
	return database.DB.Create(box).Error
}

func GetRootBoxes() ([]models.Box, error) {
	var boxes []models.Box

	// 查询所有 ParentID 为 nil 的 Box（即根节点）
	if err := database.DB.Where("parent_id IS NULL").Find(&boxes).Error; err != nil {
		return nil, err
	}

	return boxes, nil
}

// GetBoxByID 根据ID获取Box
func GetBoxByID(id int) (*models.Box, error) {
	var box models.Box
	err := database.DB.Preload("Children").Preload("Files").First(&box, id).Error
	return &box, err
}

// UpdateBox 更新Box信息
func UpdateBox(box *models.Box) error {
	return database.DB.Save(box).Error
}

// DeleteBox 删除Box
func DeleteBox(id int) error {
	return database.DB.Delete(&models.Box{}, id).Error
}

// AddBoxChild 添加子Box
func AddBoxChild(parentID int, childID int) error {
	return database.DB.Model(&models.Box{ID: parentID}).Association("Children").Append(&models.Box{ID: childID})
}

// RemoveBoxChild 移除子Box
func RemoveBoxChild(parentID int, childID int) error {
	return database.DB.Model(&models.Box{ID: parentID}).Association("Children").Delete(&models.Box{ID: childID})
}

// AddBoxFile 添加文件到Box
func AddBoxFile(boxID int, fileID int) error {
	return database.DB.Model(&models.Box{ID: boxID}).Association("Files").Append(&models.File{ID: fileID})
}

// RemoveBoxFile 从Box移除文件
func RemoveBoxFile(boxID int, fileID int) error {
	return database.DB.Model(&models.Box{ID: boxID}).Association("Files").Delete(&models.File{ID: fileID})
}

// GetBoxByName 根据名称获取Box
func GetBoxByName(name string) (*models.Box, error) {
	var box models.Box
	err := database.DB.Where("name = ?", name).First(&box).Error
	return &box, err
}
