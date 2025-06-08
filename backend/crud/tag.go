package crud

import (
	"box/backend/database"
	"box/backend/models"
	"database/sql"
)

// CreateTag 创建标签记录
func CreateTag(tag *models.Tag) error {
	return database.DB.Create(tag).Error
}

// GetTagByID 根据ID获取标签
func GetTagByID(id int) (*models.Tag, error) {
	var tag models.Tag
	err := database.DB.First(&tag, id).Error
	return &tag, err
}

// UpdateTag 更新标签信息
func UpdateTag(tag *models.Tag) error {
	return database.DB.Save(tag).Error
}

// DeleteTag 删除标签
func DeleteTag(id int) error {
	return database.DB.Delete(&models.Tag{}, id).Error
}

// SetTagColor 设置标签颜色
func SetTagColor(id int, color string) error {
	return database.DB.Model(&models.Tag{}).Where("id = ?", id).Update("color", sql.NullString{String: color, Valid: true}).Error
}

// ClearTagColor 清除标签颜色
func ClearTagColor(id int) error {
	return database.DB.Model(&models.Tag{}).Where("id = ?", id).Update("color", sql.NullString{Valid: false}).Error
}
