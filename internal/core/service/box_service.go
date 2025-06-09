package service

import (
	"box/internal/storage/models"
	"box/internal/storage/repos"
	"log"
)

type BoxService interface {
	GetRootBoxes() ([]models.Box, error)
	GetBoxByID(id uint) (*models.Box, error)
	GetChildBoxes(parentID uint) ([]models.Box, error)
	GetFilesInBox(boxID uint) ([]models.File, error)
	CreateBox(name string) (*models.Box, error)
}

type LocalBoxService struct{}

func (s *LocalBoxService) GetRootBoxes() ([]models.Box, error) {
	log.Printf("[BoxService] 获取根目录Box列表")
	boxes, err := repos.GetRootBoxes()
	if err != nil {
		log.Printf("[BoxService] 获取根目录失败: %v", err)
	} else {
		log.Printf("[BoxService] 获取到 %d 个根目录", len(boxes))
	}
	return boxes, err
}

func (s *LocalBoxService) GetBoxByID(id uint) (*models.Box, error) {
	log.Printf("[BoxService] 按ID查询Box: %d", id)
	box, err := repos.GetBoxByID(int(id))
	if err != nil {
		log.Printf("[BoxService] 查询Box失败(ID:%d): %v", id, err)
	} else if box == nil {
		log.Printf("[BoxService] Box不存在(ID:%d)", id)
	}
	return box, err
}

func (s *LocalBoxService) GetChildBoxes(parentID uint) ([]models.Box, error) {
	log.Printf("[BoxService] 获取子Box列表(ParentID:%d)", parentID)
	boxes, err := repos.GetBoxesByParentID(uint(parentID))
	if err != nil {
		log.Printf("[BoxService] 获取子Box失败(ParentID:%d): %v", parentID, err)
	} else {
		log.Printf("[BoxService] 获取到 %d 个子Box", len(boxes))
	}
	return boxes, err
}

func (s *LocalBoxService) GetFilesInBox(boxID uint) ([]models.File, error) {
	log.Printf("[BoxService] 获取Box内文件列表(BoxID:%d)", boxID)
	files, err := repos.GetFilesByBoxID(boxID)
	if err != nil {
		log.Printf("[BoxService] 获取文件列表失败(BoxID:%d): %v", boxID, err)
	} else {
		log.Printf("[BoxService] 获取到 %d 个文件", len(files))
	}
	return files, err
}

func (s *LocalBoxService) CreateBox(name string) (*models.Box, error) {
	log.Printf("[BoxService] 创建新Box: %s", name)
	box := &models.Box{Name: name}
	err := repos.CreateBox(box)
	if err != nil {
		log.Printf("[BoxService] 创建Box失败: %s, 错误: %v", name, err)
	} else {
		log.Printf("[BoxService] 成功创建Box(ID:%d)", box.ID)
	}
	return box, err
}
